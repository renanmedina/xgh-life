package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SlackBotHandler(c *gin.Context) {
	logger := log.Default()
	logger.Print(c.Params)
	c.JSON(http.StatusOK, gin.H{"results": "ok"})
}
