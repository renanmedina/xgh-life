package configs

import (
	"os"
	"strconv"
)

type ApplicationConfigs struct {
	NewRelicEnabled    bool
	NewRelicAppName    string
	NewRelicLicenseKey string
	GithubAuthToken    string
}

func NewApplicationConfigs() *ApplicationConfigs {
	newRelicEnabled, err := strconv.ParseBool(os.Getenv("NEWRELIC_ENABLED"))

	if err != nil {
		newRelicEnabled = false
	}

	config := &ApplicationConfigs{
		NewRelicEnabled:    newRelicEnabled,
		NewRelicAppName:    os.Getenv("NEWRELIC_APP_NAME"),
		NewRelicLicenseKey: os.Getenv("NEWRELIC_LICENSE_KEY"),
		GithubAuthToken:    os.Getenv("GITHUB_AUTH_TOKEN"),
	}

	return config
}
