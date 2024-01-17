package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-bot/handlers"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Static("/images/", "./frontend/static/images")
	router.Static("/scripts/", "./frontend/static/scripts")
	router.LoadHTMLGlob("./frontend/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.AddParam("id", "roulette")
		handlers.AxiomDetailsHandlerHtml(c)
	})

	router.GET("/:id", handlers.AxiomDetailsHandlerHtml)
	router.GET("/axioms/:id", handlers.AxiomDetailsHandlerHtml)

	api := router.Group("/api")
	{
		api.GET("/axioms", handlers.AxiomsListHandler)
		api.GET("/axioms/:id", handlers.AxiomDetailsHandlerJson)
	}

	slack := router.Group("/slack")
	{
		slack.POST("/axioms", handlers.SlackBotHandler)
	}

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
