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
	newRelicEnabled, err := strconv.ParseBool(os.Getenv("NEW_RELIC_ENABLED"))

	if err != nil {
		newRelicEnabled = false
	}

	config := &ApplicationConfigs{
		NewRelicEnabled:    newRelicEnabled,
		NewRelicAppName:    os.Getenv("NEW_RELIC_APP_NAME"),
		NewRelicLicenseKey: os.Getenv("NEW_RELIC_LICENSE_KEY"),
		GithubAuthToken:    os.Getenv("GITHUB_AUTH_TOKEN"),
	}

	return config
}
