package integrations

import (
	"log"

	"github.com/renanmedina/xgh-bot/configs"
)

var Logger *log.Logger

func NewApplicationLogger() *log.Logger {
	if Logger != nil {
		return Logger
	}

	appConfig := configs.NewApplicationConfigs()

	Logger = log.Default()

	if appConfig.NewRelicEnabled {
		Logger.Println("Initializing and using newrelic logger agent")
		Logger = NewRelicLogger()
	}

	return Logger
}
