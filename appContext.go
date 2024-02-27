package inkstone

import (
	"embed"

	"github.com/jmoiron/sqlx"
)

type AppContextAwareFunc func(*AppContext)

type AppContext struct {
	locales *embed.FS
	*Env
	*sqlx.DB
}

func NewAppContext(locales *embed.FS) *AppContext {
	return NewAppContextWithEnv(locales, LoadEnv())
}

func NewAppContextWithEnv(locales *embed.FS, env *Env) *AppContext {
	db := ConnectDB(
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

func (app *AppContext) Close() {
	app.DB.Close()
}
