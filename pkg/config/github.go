package config

import (
	"context"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func (cfg GitHubBillingExporterConfig) GetGitHubClient() *github.Client {
	if *cfg.githubToken == "" {
		return github.NewClient(nil)
	}
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *cfg.githubToken},
	)
	httpClient := oauth2.NewClient(ctx, tokenSource)
	return github.NewClient(httpClient)
}
