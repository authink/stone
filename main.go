package inkstone

import (
	"net/http"

	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/env"
	"github.com/authink/inkstone/web"
	"github.com/gin-gonic/gin"
)

func Run(opts *app.Options) {
	appCtx := app.NewAppContext(opts.Locales)
	defer appCtx.Close()

	if appCtx.AppENV != env.DEVELOPMENT {
		gin.SetMode("release")
	}

	app.Run(
		func() http.Handler {
			return web.SetupRouterWith(appCtx, opts)
		},
		appCtx,
		opts,
	)
}
