package inkstone

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	APP_KEY = "is_app"
)

type Context struct {
	*gin.Context
}

func (c *Context) setApp(app *AppContext) {
	c.Set(APP_KEY, app)
}

func (c *Context) App() *AppContext {
	return c.MustGet(APP_KEY).(*AppContext)
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

func translateErrorMsg(c *Context, err error) {
	if e, ok := err.(*ClientError); ok {
		e.Message = Translate(c.Context, e.Code)
	}
}
