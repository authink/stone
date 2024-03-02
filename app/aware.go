package app

import "github.com/gin-gonic/gin"

type AppContextAwareFunc func(*AppContext)

type APIGroupAwareFunc func(*gin.RouterGroup)
