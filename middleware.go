package inkstone

import (
	"embed"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	libI18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func setupAppMiddleware(app *AppContext) gin.HandlerFunc {
	return HandlerAdapter(func(c *Context) {
		c.setApp(app)
		c.Next()
	})
}

func i18nBundle(locales *embed.FS) (bundle *libI18n.Bundle) {
	bundle = libI18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFileFS(locales, "locales/en.toml")
	bundle.LoadMessageFileFS(locales, "locales/zh-CN.toml")
	return
}

func setupI18nMiddleware(locales *embed.FS) gin.HandlerFunc {

	return HandlerAdapter(func(c *Context) {
		lang := c.Query("lang")
		accept := c.GetHeader("Accept-Language")
		localizer := libI18n.NewLocalizer(i18nBundle(locales), lang, accept)

		c.setLocalizer(localizer)
		c.Next()
	})
}
