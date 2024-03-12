package stone

import (
	"net/http"

	"github.com/authink/stone/app"
	"github.com/authink/stone/env"
	"github.com/authink/stone/web"
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
