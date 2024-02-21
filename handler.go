package inkstone

import "github.com/gin-gonic/gin"

type HandlerFunc func(*Context)

func HandlerAdapter(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&Context{c})
	}
}
