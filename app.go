package inkstone

import (
	"embed"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

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
	)

	return &AppContext{locales, env, db}
}

func (app *AppContext) Close() {
	app.DB.Close()
}

func setupAppMiddleware(app *AppContext) gin.HandlerFunc {
	return HandlerAdapter(func(c *Context) {
		c.setApp(app)
		c.Next()
	})
}
