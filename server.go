package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-bot/handlers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	port := os.Getenv("PORT")
	logger := log.Default()

	if port == "" {
		logger.Println("[WEBSERVER-LOG] No provided port by environment, using default 8080")
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

	host := fmt.Sprintf("localhost:%s", port)
	logger.Printf("[WEBSERVER-LOG] Listening and serving HTTP on %s", host)

	err := router.Run(host)
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
