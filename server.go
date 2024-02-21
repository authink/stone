package inkstone

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

func gracefulShutdown(app *AppContext, srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(app.Env.ShutdownTimeout)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
}
