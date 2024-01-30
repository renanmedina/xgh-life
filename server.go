package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-bot/handlers"
)

var (
	//go:embed frontend/static/*
	staticEmbeded embed.FS

	//go:embed frontend/templates/*
	templatesEmbeded embed.FS
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	port := detectEnvPort()
	logger := log.Default()

	if port == "" {
		logger.Println("[XGH-BOT:WEBSERVER-LOG] No provided port by environment, using default 8080")
		port = "8080"
	}

	assetsFS, _ := fs.Sub(staticEmbeded, "frontend")
	router.StaticFS("/assets", http.FS(assetsFS))

	templ := template.Must(template.New("").ParseFS(templatesEmbeded, "frontend/templates/*.tmpl"))
	router.SetHTMLTemplate(templ)

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

	host := fmt.Sprintf(":%s", port)
	logger.Printf("[XGH-BOT:WEBSERVER-LOG] Listening and serving HTTP on %s", host)

	err := router.Run(host)
	if err != nil {
		panic("[XGH-BOT:WEBSERVER-LOG] failed to start server due to: " + err.Error())
	}
}

func detectEnvPort() string {
	port := os.Getenv("PORT")

	if port != "" {
		return port
	}

	return os.Getenv("XGH_HTTP_PORT")
}
