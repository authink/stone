package app

import (
	"embed"

	"github.com/authink/inkstone/db"
	"github.com/authink/inkstone/env"
)

type AppContext struct {
	Locales *embed.FS
	*env.Env
	*db.DB
}

func NewAppContext(locales *embed.FS) *AppContext {
	return NewAppContextWithEnv(locales, env.LoadEnv())
}

func NewAppContextWithEnv(locales *embed.FS, env *env.Env) *AppContext {
	db := db.ConnectDB(
		env.DbUser,
		env.DbPasswd,
		env.DbName,
		env.DbHost,
		env.DbPort,
		env.DbMaxOpenConns,
		env.DbMaxIdleConns,
		env.DbConnMaxLifeTime,
		env.DbConnMaxIdleTime,
		env.DbLogMode,
	)

	return &AppContext{locales, env, db}
}

func (appCtx *AppContext) Close() {
	appCtx.DB.Close()
}
