package integrations

import (
	"log"

	"github.com/renanmedina/xgh-bot/configs"
)

var Logger *log.Logger

func init() {

}

func NewApplicationLogger() *log.Logger {
	if Logger != nil {
		return Logger
	}

	appConfig := configs.NewApplicationConfigs()

	if appConfig.NewRelicEnabled {
		return NewRelicLogger()
	}

	return log.Default()
}
