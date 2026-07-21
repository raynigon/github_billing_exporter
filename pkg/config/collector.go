package config

import (
	"strings"

	"github.com/raynigon/github_billing_exporter/v2/collector"
)

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func (cfg GitHubBillingExporterConfig) GetCollectorConfig() collector.CollectorConfig {
	orgs := strings.Split(cfg.GithubOrgs, " ")
	orgs = filter(orgs, func(item string) bool { return item != "" })
	disabledCollectors := strings.Split(cfg.DisabledCollectors, " ")
	disabledCollectors = filter(disabledCollectors, func(item string) bool { return item != "" })
	return collector.CollectorConfig{
		Logger:             cfg.GetLogger(),
		Github:             cfg.GetGitHubClient(),
		Orgs:               orgs,
		DisabledCollectors: disabledCollectors,
	}
}
