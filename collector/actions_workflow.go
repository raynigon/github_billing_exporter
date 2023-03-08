package collector

import (
	"context"
	"strconv"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/raynigon/github_billing_exporter/v2/pkg/gh_workflow"
)

var (
	repoActionsSubsystem = "actions_workflow"
)

type WorkflowActionsCollector struct {
	config          CollectorConfig
	metrics         map[string]*gh_workflow.GitHubWorkflowMetrics
	usedMinutesReal *prometheus.Desc
}

func init() {
	registerCollector(repoActionsSubsystem, NewWorkflowActionsCollector)
}

// NewRepoActionsCollector returns a new Collector exposing actions billing stats.
func NewWorkflowActionsCollector(config CollectorConfig, ctx context.Context) (Collector, error) {
	collector := &WorkflowActionsCollector{
		config:  config,
		metrics: make(map[string]*gh_workflow.GitHubWorkflowMetrics),
		usedMinutesReal: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, repoActionsSubsystem, "minutes_real_count"),
			"GitHub actions used minutes without platform multiplier",
			[]string{"org", "repository", "workflow_name", "workflow_id", "type"}, nil,
		),
	}
	err := collector.Reload(ctx)
	if err != nil {
		return nil, err
	}
	return collector, nil
}

// Describe implements Collector.
func (wac *WorkflowActionsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- wac.usedMinutesReal
}

func (wac *WorkflowActionsCollector) Reload(ctx context.Context) error {
	metrics := make(map[string]*gh_workflow.GitHubWorkflowMetrics)
	for _, org := range wac.config.Orgs {
		metrics[org] = gh_workflow.NewGitHubWorkflowMetrics(wac.config.Github, org, ctx)
	}
	wac.metrics = metrics
	return nil
}

func (wac *WorkflowActionsCollector) Update(ctx context.Context, ch chan<- prometheus.Metric) error {
	wg := sync.WaitGroup{}
	wg.Add(len(wac.metrics))
	for org, repoMetrics := range wac.metrics {
		go func(org string, repoMetrics *gh_workflow.GitHubWorkflowMetrics) {
			for _, workflow := range repoMetrics.CollectActions(ctx) {
				if workflow.Usage.Billable == nil {
					continue
				}
				repoName := *workflow.Repository.Name
				workflowName := *workflow.Workflow.Name
				workflowId := strconv.FormatInt(*workflow.Workflow.ID, 10)
				for name, value := range *workflow.Usage.Billable {
					if value.TotalMS == nil {
						continue
					}
					ch <- prometheus.MustNewConstMetric(wac.usedMinutesReal, prometheus.CounterValue, float64(*value.TotalMS)/60_000.0, org, repoName, workflowName, workflowId, name)
				}
			}
			wg.Done()
		}(org, repoMetrics)
	}
	wg.Wait()
	return nil
}
