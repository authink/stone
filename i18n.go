package inkstone

import (
	"embed"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	libI18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func newI18nBundle(locales *embed.FS) (bundle *libI18n.Bundle) {
	bundle = libI18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFileFS(locales, "locales/en.toml")
	bundle.LoadMessageFileFS(locales, "locales/zh-CN.toml")
	return
}

func setupI18nMiddleware(locales *embed.FS) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.Query("lang")
		accept := c.GetHeader("Accept-Language")
		localizer := libI18n.NewLocalizer(newI18nBundle(locales), lang, accept)

		c.Set("localizer", localizer)
		c.Next()
	}
}

func Translate(c *gin.Context, messageID string) string {
	localizer := c.MustGet("localizer").(*libI18n.Localizer)
	return localizer.MustLocalize(&libI18n.LocalizeConfig{MessageID: messageID})
}
