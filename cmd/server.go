package main

import (
	"fmt"
	"net/http"
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

	router.Static("/_app/", "./frontend/build/_app")
	router.Static("/images/", "./frontend/build/images")
	router.LoadHTMLFiles("frontend/build/index.html")

	router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	api := router.Group("/api")
	{
		api.GET("/axioms", handlers.AxiomsListHandler)
		api.GET("/axioms/:id", handlers.AxiomDetailsHandler)
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
