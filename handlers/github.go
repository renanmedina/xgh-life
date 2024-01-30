package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GithubWebhookMessage struct {
	eventType string
	Action    string `json:"action" binding:"required"`
}

func GithubAutoApprovePRHandler(c *gin.Context) {
	var message GithubWebhookMessage
	c.Bind(message)
	message.eventType = c.GetHeader("X-Github-Event")

	fmt.Println(message)

	// if message.isPullRequestOpened() {
	// 	appConfigs, err := configs.NewApplicationConfigs()

	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	client := integrations.NewGithubClient(appConfigs.GithubAuthToken)
	// 	// client.ApprovePullRequest(message.repositoryNameWithOwner(), message.pullRequestId())
	// }

	c.JSON(http.StatusAccepted, gin.H{"message": "Accepted"})
}

func (m *GithubWebhookMessage) isPullRequestOpened() bool {
	return m.Action == "opened"
}

func (m *GithubWebhookMessage) repositoryNameWithOwner() string {
	return ""
}

func (m *GithubWebhookMessage) pullRequestId() string {
	return ""
}
