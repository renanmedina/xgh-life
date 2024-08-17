package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/renanmedina/xgh-life/gohorse"
)

func LanguageDiscovery() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("language", gohorse.DEFAULT_LANGUAGE)
	}
}
