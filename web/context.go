package web

import (
	"net/http"

	"github.com/authink/stone/app"
	"github.com/gin-gonic/gin"
	libI18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

const (
	APP_CONTEXT_KEY = "is_app_context"
	LOCALIZER_KEY   = "is_localizer"
)

type Context struct {
	*gin.Context
}

func (c *Context) AppContext() *app.AppContext {
	return c.MustGet(APP_CONTEXT_KEY).(*app.AppContext)
}

func (c *Context) Localizer() *libI18n.Localizer {
	return c.MustGet(LOCALIZER_KEY).(*libI18n.Localizer)
}

func (c *Context) Response(res any) {
	c.JSON(http.StatusOK, res)
}

func (c *Context) Empty() {
	c.Status(http.StatusNoContent)
}

func (c *Context) AbortWithClientError(err error) {
	translateErrorMsg(c, err)
	c.AbortWithStatusJSON(
		http.StatusBadRequest,
		err,
	)
}

func (c *Context) AbortWithUnauthorized(err error) {
	translateErrorMsg(c, err)
	c.AbortWithStatusJSON(
		http.StatusUnauthorized,
		err,
	)
}

func (c *Context) AbortWithForbidden(err error) {
	translateErrorMsg(c, err)
	c.AbortWithStatusJSON(
		http.StatusForbidden,
		err,
	)
}

func (c *Context) AbortWithServerError(err error) {
	c.AbortWithError(
		http.StatusInternalServerError,
		err,
	)
}

func (c *Context) Translate(messageID string) string {
	return c.Localizer().MustLocalize(&libI18n.LocalizeConfig{MessageID: messageID})
}

func translateErrorMsg(c *Context, err error) {
	if e, ok := err.(*ClientError); ok {
		e.Message = c.Translate(e.Code)
	}
}

func (c *Context) setAppContext(appCtx *app.AppContext) {
	c.Set(APP_CONTEXT_KEY, appCtx)
}

func (c *Context) setLocalizer(localizer *libI18n.Localizer) {
	c.Set(LOCALIZER_KEY, localizer)
}
