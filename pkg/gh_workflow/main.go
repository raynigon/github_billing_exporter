package gh_workflow

import (
	"context"
	"sync"

	"github.com/google/go-github/v50/github"
)

type GitHubWorkflowMetrics struct {
	github       *github.Client
	organization string
	workflows    []GitHubWorkflow
}

type GitHubWorkflowMeasurement struct {
	Repository *github.Repository
	Workflow   *github.Workflow
	Usage      *github.WorkflowUsage
}

func NewGitHubWorkflowMetrics(github *github.Client, org string, ctx context.Context) *GitHubWorkflowMetrics {
	workflows := listWorkflows(github, ctx, org)
	return &GitHubWorkflowMetrics{github, org, workflows}
}

func (github *GitHubWorkflowMetrics) CollectActions(ctx context.Context) []GitHubWorkflowMeasurement {
	c := make(chan GitHubWorkflowMeasurement, 100)
	wg := sync.WaitGroup{}
	wg.Add(len(github.workflows))
	for _, workflow := range github.workflows {
		go func(workflow GitHubWorkflow) {
			model, _, err := github.github.Actions.GetWorkflowUsageByID(
				ctx,
				github.organization,
				*workflow.repository.Name,
				*workflow.workflow.ID,
			)
			if err != nil {
				wg.Done()
				return
			}
			c <- GitHubWorkflowMeasurement{
				Repository: workflow.repository,
				Workflow:   workflow.workflow,
				Usage:      model,
			}
			wg.Done()
		}(workflow)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	// Collect Results
	result := []GitHubWorkflowMeasurement{}
	for measurement := range c {
		result = append(result, measurement)
	}
	return result
}
