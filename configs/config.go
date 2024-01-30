package configs

import (
	"fmt"
	"os"
	"strconv"
)

type ApplicationConfigs struct {
	NewRelicEnabled    bool
	NewRelicAppName    string
	NewRelicLicenseKey string
}

func NewApplicationConfigs() (*ApplicationConfigs, error) {
	newRelicEnabled, err := strconv.ParseBool(os.Getenv("NEWRELIC_ENABLED"))

	if err != nil {
		return nil, fmt.Errorf("[XGH-BOT:Application] Invalid value for NEWRELIC_ENABLED environment variable: %s", err.Error())
	}

	config := &ApplicationConfigs{
		NewRelicEnabled:    newRelicEnabled,
		NewRelicAppName:    os.Getenv("NEWRELIC_APP_NAME"),
		NewRelicLicenseKey: os.Getenv("NEWRELIC_LICENSE_KEY"),
	}

	return config, nil
}
