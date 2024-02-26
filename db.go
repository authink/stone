package inkstone

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/qustavo/sqlhooks/v2"
)

type hooksContextKey string

var hooksCtxAppKey = hooksContextKey("begin")

type hooks struct{}

func (h *hooks) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	return context.WithValue(ctx, hooksCtxAppKey, time.Now()), nil
}

func (h *hooks) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	begin := ctx.Value(hooksCtxAppKey).(time.Time)
	log.Printf("> %s %v. took: %s\n", query, args, time.Since(begin))
	return ctx, nil
}

func (h *hooks) OnError(ctx context.Context, err error, query string, args ...interface{}) error {
	// log.Printf("> %s %v. error: %v\n", query, args, err)
	return nil
}

func registerMysqlDriverIfNeeded(logMode bool) (mysqlDriverName string) {
	mysqlDriverName = "mysql"
	if logMode {
		mysqlDriverName = "inkMysql"
		sql.Register(mysqlDriverName, sqlhooks.Wrap(&mysql.MySQLDriver{}, &hooks{}))
	}
	return
}

func ConnectDBUrl(user, password, dbName, host string, port uint16, withSchema bool) string {
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbName)

	if withSchema {
		return "mysql://" + databaseUrl
	}

	return databaseUrl
}

func ConnectDB(user, password, dbName, host string, port, maxOpenConns, maxIdleConns, connMaxLifeTime, connMaxIdleTime uint16, logMode bool) *sqlx.DB {
	mysqlDriverName := registerMysqlDriverIfNeeded(logMode)

	databaseUrl := ConnectDBUrl(user, password, dbName, host, port, false)

	db, err := sqlx.Open(mysqlDriverName, databaseUrl)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(int(maxOpenConns))
	db.SetMaxIdleConns(int(maxIdleConns))
	db.SetConnMaxLifetime(time.Duration(connMaxLifeTime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(connMaxIdleTime) * time.Second)

	return db
}

func CreateDB(user, password, dbName, host string, port uint16) func() {
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/", user, password, host, port)

	db, err := sqlx.Open("mysql", databaseUrl)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	if err != nil {
		db.Close()
		panic(err)
	}

	return func() {
		defer db.Close()
		_, err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName))
		if err != nil {
			panic(err)
		}
	}
}
