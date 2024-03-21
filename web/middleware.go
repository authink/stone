package web

import (
	"embed"

	"github.com/BurntSushi/toml"
	"github.com/authink/stone/app"
	"github.com/gin-gonic/gin"
	libI18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func setupAppMiddleware(appCtx *app.AppContext) gin.HandlerFunc {
	return HandlerAdapter(func(c *Context) {
		c.setAppContext(appCtx)
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
		locale := c.Query("locale")
		accept := c.GetHeader("Accept-Language")
		localizer := libI18n.NewLocalizer(i18nBundle(locales), locale, accept)

		c.setLocalizer(localizer)
		c.Next()
	})
}
