package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-life/gohorse"
	"github.com/renanmedina/xgh-life/integrations"
)

type GithubWebhookMessage struct {
	eventType string
	Action    string `json:"action"`

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
	logger := integrations.NewApplicationLogger()
	messageJson, err := c.GetRawData()

	if err != nil {
		logger.Printf("[XGH-LIFE:WEBSERVER-LOG] Failed to parse github event received")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	githubEventType := c.GetHeader("X-Github-Event")
	logger.Printf("[XGH-LIFE:WEBSERVER-LOG] Github webhook event received: %s", messageJson)
	logger.Printf("[XGH-LIFE:WEBSERVER-LOG] Github webhook event_type received: %s", githubEventType)

	var message GithubWebhookMessage
	json.Unmarshal(messageJson, &message)
	message.eventType = githubEventType

	if message.isPullRequestOpened() {
		logger.Printf("[XGH-LIFE:WEBSERVER-LOG] Running pull request auto approver")
		use_case := gohorse.NewAutoApprovePullRequestUseCase(logger)
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
