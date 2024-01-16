package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-bot/gohorse"
	"github.com/renanmedina/xgh-bot/slack"
)

type CommandRequest struct {
	CommandType   string `json:"command" form:"command" binding:"required"`
	RequestedType string `json:"text" form:"text" binding:"required"`
}

func (c CommandRequest) isXGH() bool {
	return c.CommandType == "/xgh"
}

func SlackBotHandler(c *gin.Context) {
	var request CommandRequest
	c.Bind(&request)

	if !request.isXGH() {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Command %s not available", request.CommandType)})
		return
	}

	use_case := gohorse.NewGetAxiomUseCase()
	axiom, err := use_case.Execute(request.RequestedType)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slack.NewSlackSimpleResponse(slack.REPLACE_ORIGINAL, axiom.ToQuote()))
}
