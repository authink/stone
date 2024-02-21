package inkstone

import (
	"fmt"
	"net/http"
)

func startServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func createServer(app *AppContext, handler http.Handler) (srv *http.Server) {

	srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", app.Env.Host, app.Env.Port),
		Handler: handler,
	}

	go startServer(srv)

	return
}
