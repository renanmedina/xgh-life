package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-bot/handlers"
)

func main() {
	router := gin.Default()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	api := router.Group("/api")
	{
		api.GET("/axioms", handlers.AxiomsListHandler)
		api.GET("/axioms/:number", handlers.AxiomDetailsHandler)
	}

	slack := router.Group("/slack")
	{
		slack.POST("/slack/axioms", handlers.SlackBotHandler)
	}

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
