package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/qustavo/sqlhooks/v2"
)

type hooksContextKey string

var hooksCtxAppKey = hooksContextKey("beginTime")

type hooks struct{}

func (h *hooks) Before(ctx context.Context, query string, args ...any) (context.Context, error) {
	return context.WithValue(ctx, hooksCtxAppKey, time.Now()), nil
}

func (h *hooks) After(ctx context.Context, query string, args ...any) (context.Context, error) {
	begin := ctx.Value(hooksCtxAppKey).(time.Time)
	log.Printf("> %s %v. took: %s\n", query, args, time.Since(begin))
	return ctx, nil
}

func (h *hooks) OnError(ctx context.Context, err error, query string, args ...any) error {
	// log.Printf("> %s %v. error: %v\n", query, args, err)
	return nil
}

func registerMysqlDriverIfNeeded(logMode bool) (mysqlDriverName string) {
	mysqlDriverName = "mysql"
	if logMode {
		mysqlDriverName = "inkMysql"
		sql.Register(mysqlDriverName, sqlhooks.Wrap(new(mysql.MySQLDriver), new(hooks)))
	}
	return
}
