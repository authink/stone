package inkstone

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type TestContextKey string

var TestCtxAppKey = TestContextKey("app")
var TestCtxRouterKey = TestContextKey("router")

func setup(app *AppContext, seed SeedFunc) {
	migrateSchema(app, "up")
	seed(app)
}

func teardown(app *AppContext) {
	migrateSchema(app, "down")
}

func TestMain(ctx *context.Context, app *AppContext, router *gin.Engine, seed SeedFunc) func(*testing.M) {
	*ctx = context.WithValue(
		*ctx,
		TestCtxAppKey,
		app,
	)
	*ctx = context.WithValue(
		*ctx,
		TestCtxRouterKey,
		router,
	)

	return func(m *testing.M) {
		setup(app, seed)

		exitCode := m.Run()

		teardown(app)

		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}
}

func TestFetch(ctx context.Context, method, pathname string, reqObj, resObj any, accessToken string) (w *httptest.ResponseRecorder, err error) {
	app := ctx.Value(TestCtxAppKey).(*AppContext)

	var reader io.Reader
	if reqObj != nil {
		reqBody, _ := json.Marshal(reqObj)
		reader = strings.NewReader(string(reqBody))
	}

	w = httptest.NewRecorder()
	req, _ := http.NewRequest(
		method,
		path.Join("/", app.Env.BasePath, pathname),
		reader,
	)

	if accessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}

	router := ctx.Value(TestCtxRouterKey).(*gin.Engine)
	router.ServeHTTP(w, req)

	if w.Body.Len() > 0 {
		err = json.Unmarshal(w.Body.Bytes(), resObj)
	}
	return
}