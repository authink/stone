package db

import (
	"fmt"
	"net/url"
	"time"

	"github.com/jmoiron/sqlx"
)

func MakeUrl(user, password, dbName, host string, port uint16, timeZone string, withSchema bool) string {
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=%s", user, password, host, port, dbName, url.QueryEscape(timeZone))

	if withSchema {
		return "mysql://" + databaseUrl
	}

	return databaseUrl
}

func Connect(user, password, dbName, host string, port, maxOpenConns, maxIdleConns, connMaxLifeTime, connMaxIdleTime uint16, timeZone string, logMode bool) *DB {
	mysqlDriverName := registerMysqlDriverIfNeeded(logMode)

	databaseUrl := MakeUrl(user, password, dbName, host, port, timeZone, false)

	db, err := sqlx.Open(mysqlDriverName, databaseUrl)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(int(maxOpenConns))
	db.SetMaxIdleConns(int(maxIdleConns))
	db.SetConnMaxLifetime(time.Duration(connMaxLifeTime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(connMaxIdleTime) * time.Second)

	return &DB{db}
}

func CreateTestDB(user, password, dbName, host string, port uint16) func() {
	var (
		databaseUrl = fmt.Sprintf("%s:%s@tcp(%s:%d)/", user, password, host, port)

		connect = func() *sqlx.DB {
			db, err := sqlx.Open("mysql", databaseUrl)
			if err != nil {
				panic(err)
			}
			return db
		}

		create = func(db *sqlx.DB) {
			_, err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
			if err != nil {
				panic(err)
			}
		}

		drop = func(db *sqlx.DB) {
			_, err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName))
			if err != nil {
				panic(err)
			}
		}
	)

	db := connect()
	defer db.Close()
	drop(db)
	create(db)

	return func() {
		db := connect()
		defer db.Close()
		drop(db)
	}
}
