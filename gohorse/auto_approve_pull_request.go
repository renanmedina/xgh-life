package gohorse

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/renanmedina/xgh-life/configs"
	"github.com/renanmedina/xgh-life/integrations"
)

type AutoApprovePullRequest struct {
	githubClient *integrations.GithubClient
	logger       *log.Logger
	whitelist    []Author
}

type Author struct {
	name          string
	approveChance int
}

func (a *Author) shouldAutoApprove() bool {
	return rand.Intn(100) <= a.approveChance
}

func NewAutoApprovePullRequestUseCase(logger *log.Logger) *AutoApprovePullRequest {
	appConfigs := configs.NewApplicationConfigs()
	client := integrations.NewGithubClient(appConfigs.GithubAuthToken)
	return &AutoApprovePullRequest{
		githubClient: client,
		logger:       logger,
		whitelist: []Author{
			{"eduardohertz", 100},
			{"pedro-hgm", 100},
			{"thekaduu", 50},
		},
	}
}

func (use_case *AutoApprovePullRequest) Execute(repositoryName string, pullRequestId string, author string) error {
	whitelistedAuthor := use_case.getWhitelistedAuthor(author)
	if whitelistedAuthor == nil {
		msg := fmt.Sprintf("pull request author (%s) is not whitelisted for auto approve", author)
		use_case.logger.Printf("[XGH-LIFE:WEBSERVER-LOG:AUTO-APPOVE-PR] %s", msg)
		return fmt.Errorf(msg)
	}

	// client := use_case.githubClient
	// return client.ApprovePullRequest(repositoryName, pullRequestId)
	use_case.logger.Printf("[XGH-LIFE:WEBSERVER-LOG:AUTO-APPOVE-PR] Pull Request %s in %s by %s would have been approved automatically", repositoryName, pullRequestId, author)
	return nil
}

func (use_case *AutoApprovePullRequest) getWhitelistedAuthor(authorName string) *Author {
	for _, whitelistAuthor := range use_case.whitelist {
		if whitelistAuthor.name == authorName && whitelistAuthor.shouldAutoApprove() {
			return &whitelistAuthor
		}
	}

	return nil
}
