package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-life/gohorse"
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
	hostname := request.Host

	splits := strings.Split(hostname, ".")
	subdomainLang := splits[0]

	// expecting domain to be something like en.xgh.life or en.staging.xgh.life
	if len(splits) > 2 && len(subdomainLang) == 2 {
		return prefixDomainToLang(subdomainLang)
	}

	return ""
}

func checkByIpGeolocation(context *gin.Context) string {
	// disable for now
	// userIp := context.ClientIP()
	// lookupService := integrations.NewIPGeolocationService(utils.GetConfigs().IPGeolocationApiToken)

	// locationData, err := lookupService.Lookup(userIp)

	// if err != nil {
	// 	return ""
	// }

	// if locationData.CountryCode != "BR" {
	// 	return prefixDomainToLang("en")
	// }

	return ""
}

func prefixDomainToLang(prefix string) string {
	prefixes := map[string]string{
		"en": "en-US",
		"pt": "pt-BR",
	}

	lang, found := prefixes[prefix]

	if found {
		return lang
	}

	return ""
}
