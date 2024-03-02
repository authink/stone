package test

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

	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/db"
	"github.com/authink/inkstone/env"
	"github.com/authink/inkstone/migrate"
	"github.com/authink/inkstone/web"
	"github.com/gin-gonic/gin"
)

type testContextKey string

var testCtxAppKey = testContextKey("app_context")
var testCtxRouterKey = testContextKey("router")

func setup(appCtx *app.AppContext, opts *app.Options) {
	migrate.Schema(
		"up",
		appCtx.DbMigrateFileSource,
		appCtx.DbUser,
		appCtx.DbPasswd,
		appCtx.DbName,
		appCtx.DbHost,
		appCtx.DbPort,
	)
	if opts.Seed != nil {
		opts.Seed(appCtx)
	}
}

func teardown(appCtx *app.AppContext) {
	migrate.Schema("down",
		appCtx.DbMigrateFileSource,
		appCtx.DbUser,
		appCtx.DbPasswd,
		appCtx.DbName,
		appCtx.DbHost,
		appCtx.DbPort,
	)
}

func Run(packageName string, ctx *context.Context, opts *app.Options) func(*testing.M) {
	env := env.LoadEnv()
	env.DbName = fmt.Sprintf("%s_%s", env.DbName, packageName)
	dropDB := db.CreateTestDB(
		env.DbUser,
		env.DbPasswd,
		env.DbName,
		env.DbHost,
		env.DbPort,
	)

	appCtx := app.NewAppContextWithEnv(opts.Locales, env)
	setup(appCtx, opts)
	router := web.SetupRouterWith(appCtx, opts)

	*ctx = context.WithValue(
		*ctx,
		testCtxAppKey,
		appCtx,
	)
	*ctx = context.WithValue(
		*ctx,
		testCtxRouterKey,
		router,
	)

	return func(m *testing.M) {
		defer dropDB()
		defer appCtx.Close()

		if exitCode := m.Run(); exitCode != 0 {
			os.Exit(exitCode)
		}

		teardown(appCtx)
	}
}

func Fetch(ctx context.Context, method, pathname string, reqObj, resObj any, accessToken string) (w *httptest.ResponseRecorder, err error) {
	appCtx := ctx.Value(testCtxAppKey).(*app.AppContext)

	var reader io.Reader
	if reqObj != nil {
		reqBody, _ := json.Marshal(reqObj)
		reader = strings.NewReader(string(reqBody))
	}

	w = httptest.NewRecorder()
	req, _ := http.NewRequest(
		method,
		path.Join("/", appCtx.BasePath, pathname),
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