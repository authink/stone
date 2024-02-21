package inkstone

import (
	"context"
	"embed"
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

type testContextKey string

var testCtxAppKey = testContextKey("app")
var testCtxRouterKey = testContextKey("router")

func setup(app *AppContext, seed SeedFunc) {
	migrateSchema(app, "up")
	seed(app)
}

func teardown(app *AppContext) {
	migrateSchema(app, "down")
}

func TestMain(packageName string, ctx *context.Context, locales *embed.FS, seed SeedFunc, setupAPIGroup SetupAPIGroupFunc) func(*testing.M) {
	env := LoadEnv()
	env.DbName = fmt.Sprintf("%s_%s", env.DbName, packageName)
	dropDB := CreateDB(
		env.DbUser,
		env.DbPasswd,
		env.DbName,
		env.DbHost,
		env.DbPort,
	)

	app := NewAppContextWithEnv(locales, env)
	router, apiGroup := SetupRouter(app)
	setupAPIGroup(apiGroup)

	*ctx = context.WithValue(
		*ctx,
		testCtxAppKey,
		app,
	)
	*ctx = context.WithValue(
		*ctx,
		testCtxRouterKey,
		router,
	)

	return func(m *testing.M) {
		defer dropDB()
		defer app.Close()

		setup(app, seed)

		exitCode := m.Run()

		teardown(app)

		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}
}

func TestFetch(ctx context.Context, method, pathname string, reqObj, resObj any, accessToken string) (w *httptest.ResponseRecorder, err error) {
	app := ctx.Value(testCtxAppKey).(*AppContext)

	var reader io.Reader
	if reqObj != nil {
		reqBody, _ := json.Marshal(reqObj)
		reader = strings.NewReader(string(reqBody))
	}

	w = httptest.NewRecorder()
	req, _ := http.NewRequest(
		method,
		path.Join("/", app.BasePath, pathname),
		reader,
	)

	if accessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}

	router := ctx.Value(testCtxRouterKey).(*gin.Engine)
	router.ServeHTTP(w, req)

	if w.Body.Len() > 0 {
		err = json.Unmarshal(w.Body.Bytes(), resObj)
	}
	return
}
