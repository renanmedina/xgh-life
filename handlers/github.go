package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-bot/gohorse"
)

type GithubWebhookMessage struct {
	eventType  string
	Action     string `json:"action"`
	Repository struct {
		RepositoryNameWithOwner string `json:"full_name"`
	} `json:"repository"`
	PullRequest struct {
		PullRequestId int    `json:"number"`
		State         string `json:"state"`
		Author        struct {
			Name string `json:"login"`
		} `json:"user"`
	} `json:"pull_request"`
}

func GithubAutoApprovePullRequestHandler(c *gin.Context) {
	var message GithubWebhookMessage
	c.Bind(&message)
	message.eventType = c.GetHeader("X-Github-Event")

	if message.isPullRequestOpened() {
		use_case := gohorse.NewAutoApprovePullRequestUseCase()
		err := use_case.Execute(message.repositoryNameWithOwner(), message.pullRequestId(), message.authorName())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Accepted"})
}

func (m *GithubWebhookMessage) isPullRequestOpened() bool {
	return m.eventType == "pull_request" && m.Action == "opened"
}

func (m *GithubWebhookMessage) repositoryNameWithOwner() string {
	return m.Repository.RepositoryNameWithOwner
}

func (m *GithubWebhookMessage) pullRequestId() string {
	return strconv.Itoa(m.PullRequest.PullRequestId)
}

func (m *GithubWebhookMessage) authorName() string {
	return m.PullRequest.Author.Name
}
