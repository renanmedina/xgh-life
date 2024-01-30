package integrations

import (
	"github.com/newrelic/go-agent/v3/newrelic"
)

func InitializeNewRelicApp(appName string, licenseKey string) (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
}
