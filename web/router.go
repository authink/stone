package web

import (
	"github.com/authink/stone/app"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func SetupRouter(appCtx *app.AppContext) (router *gin.Engine, gApi *gin.RouterGroup) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation(VALIDATION_EMAIL, ValidationEmail)
		v.RegisterValidation(VALIDATION_PHONE, ValidationPhone)
	}

	router = gin.Default()

	router.Use(
		setupAppMiddleware(appCtx),
		setupI18nMiddleware(appCtx.Locales),
	)

	setupSwagger(router)

	gApi = router.Group(appCtx.BasePath)
	return
}

func SetupRouterWith(appCtx *app.AppContext, opts *app.Options) *gin.Engine {
	router, apiGroup := SetupRouter(appCtx)
	if opts.SetupApiGroup != nil {
		opts.SetupApiGroup(apiGroup)
	}
	if opts.FinishSetup != nil {
		opts.FinishSetup(appCtx)
	}
	return router
}
