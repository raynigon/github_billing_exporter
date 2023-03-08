package gh_org

import (
	"context"

	"github.com/google/go-github/v50/github"
)

type GitHubOrgMetrics struct {
	github       *github.Client
	organization string
}

func NewGitHubOrgMetrics(github *github.Client, org string) *GitHubOrgMetrics {
	return &GitHubOrgMetrics{github, org}
}

func (github *GitHubOrgMetrics) CollectActions(ctx context.Context) (*github.ActionBilling, error) {
	model, _, err := github.github.Billing.GetActionsBillingOrg(ctx, github.organization)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (github *GitHubOrgMetrics) CollectPackages(ctx context.Context) (*github.PackageBilling, error) {
	model, _, err := github.github.Billing.GetPackagesBillingOrg(ctx, github.organization)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (github *GitHubOrgMetrics) CollectStorage(ctx context.Context) (*github.StorageBilling, error) {
	model, _, err := github.github.Billing.GetStorageBillingOrg(ctx, github.organization)
	if err != nil {
		return nil, err
	}
	return model, nil
}
