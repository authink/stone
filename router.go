package inkstone

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func SetupRouter(app *AppContext) (router *gin.Engine, gApi *gin.RouterGroup) {
	router = gin.Default()

	router.Use(
		setupAppMiddleware(app),
		setupI18nMiddleware(app.locales),
		setupValidationMiddleware,
	)

	setupSwagger(router)

	gApi = router.Group(app.BasePath)
	return
}

type SetupAPIGroupFunc func(*gin.RouterGroup)
