package gh_workflow

import (
	"context"

	"github.com/google/go-github/v50/github"
)

type GitHubWorkflow struct {
	repository *github.Repository
	workflow   *github.Workflow
}

func listWorkflows(client *github.Client, ctx context.Context, org string) []GitHubWorkflow {
	result := make(map[string]GitHubWorkflow)

	repos := listRepositoriesInOrg(client, org, ctx)
	for repo := range repos {
		workflows, err := listWorkflowsInRepository(client, org, *repo.Name, ctx)
		if err != nil {
			continue
		}
		for _, workflow := range workflows {
			key := org + "/" + *repo.Name + "/" + *workflow.Name
			result[key] = GitHubWorkflow{
				repository: repo,
				workflow:   workflow,
			}
		}
	}
	values := make([]GitHubWorkflow, 0, len(result))
	for _, v := range result {
		values = append(values, v)
	}
	return values
}

func listRepositoriesInOrg(client *github.Client, org string, ctx context.Context) chan *github.Repository {
	c := make(chan *github.Repository, 100)
	go func() {
		page := 1
		for {
			options := github.RepositoryListByOrgOptions{
				Type: "all",
				ListOptions: github.ListOptions{
					Page:    page,
					PerPage: 100,
				},
			}
			repos, _, err := client.Repositories.ListByOrg(ctx, org, &options)
			if err != nil {
				break
			}
			if len(repos) == 0 {
				break
			}
			for _, repo := range repos {
				c <- repo
			}
			page++
		}
		close(c)
	}()
	return c
}

func listWorkflowsInRepository(client *github.Client, org string, repo string, ctx context.Context) ([]*github.Workflow, error) {
	options := github.ListOptions{
		Page:    0,
		PerPage: 100,
	}
	workflows, _, err := client.Actions.ListWorkflows(ctx, org, repo, &options)
	if err != nil {
		return nil, err
	}
	// TODO handle if more than 100 workflows
	return workflows.Workflows, nil
}
