package gohorse

import (
	"fmt"

	"github.com/renanmedina/xgh-bot/configs"
	"github.com/renanmedina/xgh-bot/integrations"
)

type AutoApprovePullRequest struct {
	githubClient *integrations.GithubClient
	whitelist    []Author
}

type Author struct {
	name string
}

func (a *Author) shouldAutoApprove() bool {
	return true
}

func NewAutoApprovePullRequestUseCase() *AutoApprovePullRequest {
	appConfigs := configs.NewApplicationConfigs()
	client := integrations.NewGithubClient(appConfigs.GithubAuthToken)
	return &AutoApprovePullRequest{
		githubClient: client,
		whitelist: []Author{
			{"eduardohertz"},
			{"pedro-hgm"},
			{"thekaduu"},
		},
	}
}

func (use_case *AutoApprovePullRequest) Execute(repositoryName string, pullRequestId string, author string) error {
	whitelistedAuthor := use_case.getWhitelistedAuthor(author)
	if whitelistedAuthor == nil {
		return fmt.Errorf("pull request author (%s) is not whitelisted for auto approve", author)
	}

	client := use_case.githubClient
	return client.ApprovePullRequest(repositoryName, pullRequestId)
}

func (use_case *AutoApprovePullRequest) getWhitelistedAuthor(authorName string) *Author {
	for _, whitelistAuthor := range use_case.whitelist {
		if whitelistAuthor.name == authorName && whitelistAuthor.shouldAutoApprove() {
			return &whitelistAuthor
		}
	}

	return nil
}
