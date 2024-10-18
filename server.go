package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	apidocs "github.com/renanmedina/xgh-life/docs"
	"github.com/renanmedina/xgh-life/handlers"
	"github.com/renanmedina/xgh-life/integrations"
	"github.com/renanmedina/xgh-life/middlewares"
	"github.com/renanmedina/xgh-life/utils"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var (
	//go:embed frontend/static/*
	staticEmbeded embed.FS
	//go:embed frontend/templates/*
	templatesEmbeded embed.FS
)

// @title XGH Life - API Docs
// @version 1.0.0
// @contact.name Renan Medina
// @host https://xgh.life
// @BasePath /
func main() {
	gin.SetMode(gin.ReleaseMode)
	appConfigs := utils.NewApplicationConfigs()

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(middlewares.LanguageDiscovery())

	configureNewRelic(router, appConfigs)
	configureStatic(router)
	configureHandlers(router)

	err := bootServer(router)

	if err != nil {
		panic("[XGH-LIFE:WEBSERVER-LOG] failed to start server due to: " + err.Error())
	}
}

func configureNewRelic(router *gin.Engine, configs *utils.ApplicationConfigs) {
	if configs.NewRelicEnabled {
		newRelicApp, err := integrations.NewRelicApp()
		if err != nil {
			panic(fmt.Sprintf("[XGH-LIFE:NEWRELIC-INTEGRATION] %s", err))
		}

		router.Use(nrgin.Middleware(newRelicApp))
	}
}

func detectEnvPort() string {
	port := os.Getenv("PORT")

	if port != "" {
		return port
	}

	return os.Getenv("XGH_HTTP_PORT")
}

func configureStatic(router *gin.Engine) {
	assetsFS, _ := fs.Sub(staticEmbeded, "frontend")
	router.StaticFS("/assets", http.FS(assetsFS))

	templ := template.Must(template.New("").ParseFS(templatesEmbeded, "frontend/templates/*.tmpl"))
	router.SetHTMLTemplate(templ)
}

func configureHandlers(router *gin.Engine) {
	router.GET("/infra/healthcheck", handlers.HealthCheck)

	router.GET("/", func(c *gin.Context) {
		c.AddParam("id", "roulette")
		handlers.AxiomDetailsHandlerHtml(c)
	})

	router.StaticFileFS("/robots.txt", "frontend/static/robots.txt", http.FS(staticEmbeded))

	router.GET("/:id", handlers.AxiomDetailsHandlerHtml)
	router.GET("/axioms/:id", handlers.AxiomDetailsHandlerHtml)

	api := router.Group("/api")
	{
		api.GET("/random", handlers.RandomAxiomHandlerJson)
		api.GET("/axioms", handlers.AxiomsListHandler)
		api.GET("/axioms/:id", handlers.AxiomDetailsHandlerJson)

		apidocs.SwaggerInfo.BasePath = "/api/"
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	github := router.Group("/github/pull-requests")
	{
		github.POST("/auto-approve", handlers.GithubAutoApprovePullRequestHandler)
	}
}

func bootServer(router *gin.Engine) error {
	port := detectEnvPort()
	logger := integrations.NewApplicationLogger()

	if port == "" {
		logger.Println("[XGH-LIFE:WEBSERVER-LOG] No provided port by environment, using default 8080")
		port = "8080"
	}

	host := fmt.Sprintf(":%s", port)
	logger.Printf("[XGH-LIFE:WEBSERVER-LOG] Listening and serving HTTP on %s", host)

	return router.Run(host)
}
