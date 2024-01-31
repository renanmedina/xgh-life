package integrations

import (
	"log"
	"os"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/logWriter"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/renanmedina/xgh-bot/configs"
)

var appConfigs = configs.NewApplicationConfigs()
var newRelicApp *newrelic.Application

func NewRelicApp() (*newrelic.Application, error) {
	if newRelicApp != nil {
		return newRelicApp, nil
	}

	return InitializeNewRelicApp(
		appConfigs.NewRelicAppName,
		appConfigs.NewRelicLicenseKey,
	)
}

func InitializeNewRelicApp(appName string, licenseKey string) (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
}

func NewRelicLogger() *log.Logger {
	app, _ := NewRelicApp()
	writer := logWriter.New(os.Stdout, app)
	return log.New(&writer, "", log.Default().Flags())
}
