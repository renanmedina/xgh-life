package integrations

import (
	"log"

	"github.com/renanmedina/xgh-life/configs"
)

var Logger *log.Logger

func NewApplicationLogger() *log.Logger {
	if Logger != nil {
		return Logger
	}

	appConfig := configs.NewApplicationConfigs()

	Logger = log.Default()

	if appConfig.NewRelicEnabled {
		Logger = NewRelicLogger()
	}

	return Logger
}
