package utils

import (
	"os"
	"strconv"
)

type ApplicationConfigs struct {
	NewRelicEnabled       bool
	NewRelicAppName       string
	NewRelicLicenseKey    string
	GithubAuthToken       string
	IPGeolocationApiToken string
}

var loadedConfigs *ApplicationConfigs

func init() {
	loadedConfigs = NewApplicationConfigs()
}

func GetConfigs() *ApplicationConfigs {
	return loadedConfigs
}

func NewApplicationConfigs() *ApplicationConfigs {
	newRelicEnabled, err := strconv.ParseBool(os.Getenv("NEWRELIC_ENABLED"))

	if err != nil {
		newRelicEnabled = false
	}

	config := &ApplicationConfigs{
		NewRelicEnabled:       newRelicEnabled,
		NewRelicAppName:       os.Getenv("NEWRELIC_APP_NAME"),
		NewRelicLicenseKey:    os.Getenv("NEWRELIC_LICENSE_KEY"),
		GithubAuthToken:       os.Getenv("GITHUB_AUTH_TOKEN"),
		IPGeolocationApiToken: os.Getenv("IPGEOLOCATION_API_TOKEN"),
	}

	return config
}
