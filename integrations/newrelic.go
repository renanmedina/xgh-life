package integrations

import (
	"log"
	"os"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/logWriter"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/renanmedina/xgh-life/configs"
)

var appConfigs = configs.NewApplicationConfigs()
var newRelicApp *newrelic.Application

func NewRelicApp() (*newrelic.Application, error) {
	if newRelicApp != nil {
		return newRelicApp, nil
	}

	newRelicApp, err := InitializeNewRelicApp(
		appConfigs.NewRelicAppName,
		appConfigs.NewRelicLicenseKey,
		appConfigs.NewRelicEnabled,
	)

	return newRelicApp, err
}

func InitializeNewRelicApp(appName string, licenseKey string, enabled bool) (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
		newrelic.ConfigAppLogForwardingEnabled(enabled),
		newrelic.ConfigAppLogEnabled(enabled),
		newrelic.ConfigAppLogMetricsEnabled(enabled),
		newrelic.ConfigModuleDependencyMetricsEnabled(enabled),
		func(config *newrelic.Config) {
			config.Enabled = enabled
		},
	)
}

func NewRelicLogger() *log.Logger {
	app, _ := NewRelicApp()
	writer := logWriter.New(os.Stdout, app)
	writer.DebugLogging(true)
	return log.New(&writer, "", log.Default().Flags())
}
