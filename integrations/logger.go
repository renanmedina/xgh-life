package integrations

import (
	"log"

	"github.com/renanmedina/xgh-life/utils"
)

var Logger *log.Logger

func NewApplicationLogger() *log.Logger {
	if Logger != nil {
		return Logger
	}

	appConfig := utils.NewApplicationConfigs()

	Logger = log.Default()

	if appConfig.NewRelicEnabled {
		Logger = NewRelicLogger()
	}

	return Logger
}
