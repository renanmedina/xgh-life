package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-life/gohorse"
	"github.com/renanmedina/xgh-life/integrations"
	"github.com/renanmedina/xgh-life/utils"
)

func LanguageDiscovery() gin.HandlerFunc {
	return func(context *gin.Context) {
		paramLanguage := checkByParams(context)

		if paramLanguage != "" {
			context.Set("language", paramLanguage)
			return
		}

		domainLanguage := checkByHostDomain(context)

		if domainLanguage != "" {
			context.Set("language", domainLanguage)
			return
		}

		locationLanguage := checkByIpGeolocation(context)

		if locationLanguage != "" {
			context.Set("language", locationLanguage)
			return
		}

		context.Set("language", gohorse.DEFAULT_LANGUAGE)
	}
}

func checkByParams(context *gin.Context) string {
	langParam := context.Query("lang")

	if langParam != "" && gohorse.IsLanguageSupported(langParam) {
		return langParam
	}

	langParam = context.Query("language")

	if langParam != "" && gohorse.IsLanguageSupported(langParam) {
		return langParam
	}

	return ""
}

func checkByHostDomain(context *gin.Context) string {
	request := context.Request
	hostname := request.URL.Hostname()

	fmt.Println(hostname)

	return ""
}

func checkByIpGeolocation(context *gin.Context) string {
	userIp := context.ClientIP()
	lookupService := integrations.NewIPGeolocationService(utils.GetConfigs().IPGeolocationApiToken)

	locationData, err := lookupService.Lookup(userIp)

	if err != nil {
		return ""
	}

	if locationData.CountryCode != "BR" {
		return "en-US"
	}

	return ""
}
