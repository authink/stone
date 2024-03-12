package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"testing"

	"github.com/authink/stone/app"
	"github.com/authink/stone/db"
	"github.com/authink/stone/env"
	"github.com/authink/stone/migrate"
	"github.com/authink/stone/web"
	"github.com/gin-gonic/gin"
)

type testContextKey string

var testCtxAppKey = testContextKey("app_context")
var testCtxRouterKey = testContextKey("router")

func setup(env *env.Env, opts *app.Options) (appCtx *app.AppContext) {
	appCtx = app.NewAppContextWithEnv(opts.Locales, env)
	migrate.Schema(
		"up",
		appCtx.DbMigrateFileSource,
		appCtx.DbUser,
		appCtx.DbPasswd,
		appCtx.DbName,
		appCtx.DbHost,
		appCtx.DbPort,
		appCtx.DbTimeZone,
	)
	if opts.Seed != nil {
		opts.Seed(appCtx)
	}
	return
}

func teardown(appCtx *app.AppContext) {
	migrate.Schema("down",
		appCtx.DbMigrateFileSource,
		appCtx.DbUser,
		appCtx.DbPasswd,
		appCtx.DbName,
		appCtx.DbHost,
		appCtx.DbPort,
		appCtx.DbTimeZone,
	)
	appCtx.Close()
}

func Run(packageName string, ctx *context.Context, opts *app.Options) func(*testing.M) {
	env := env.Load()
	env.DbName = fmt.Sprintf("%s_%s", env.DbName, packageName)
	dropDB := db.CreateTestDB(
		env.DbUser,
		env.DbPasswd,
		env.DbName,
		env.DbHost,
		env.DbPort,
	)

	appCtx := setup(env, opts)
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
		defer teardown(appCtx)

		if exitCode := m.Run(); exitCode != 0 {
			log.Printf("Tests finished(exitCode: %d)", exitCode)
		}
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
