package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-bot/gohorse"
	"github.com/renanmedina/xgh-bot/integrations"
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

	logger := integrations.NewApplicationLogger()
	messageJson, err := parseMessage(c.Request.Body)

	if err != nil {
		logger.Printf("[XGH-BOT:WEBSERVER-LOG] Failed to parse github event received: %s", string(messageJson))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.Printf("[XGH-BOT:WEBSERVER-LOG] Github webhook event received: %s", string(messageJson))
	logger.Printf("[XGH-BOT:WEBSERVER-LOG] Github webhook event_type received: %s", message.eventType)

	if message.isPullRequestOpened() {
		logger.Printf("[XGH-BOT:WEBSERVER-LOG] Running pull request auto approver")
		use_case := gohorse.NewAutoApprovePullRequestUseCase(logger)
		err := use_case.Execute(message.repositoryNameWithOwner(), message.pullRequestId(), message.authorName())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Accepted"})
}

func parseMessage(body io.ReadCloser) (string, error) {
	bodyData, err := io.ReadAll(body)

	if err != nil {
		return "", err
	}

	messageJson, err := json.Marshal(bodyData)
	if err != nil {
		return "", err
	}

	return string(messageJson), nil
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
