package inkstone

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

func ConnectDBUrl(user, password, dbName, host string, port uint16, withSchema bool) string {
	databaseUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, port, dbName)

	if withSchema {
		return "mysql://" + databaseUrl
	}

	return databaseUrl
}

func ConnectDB(user, password, dbName, host string, port, maxOpenConns, maxIdleConns, connMaxLifeTime, connMaxIdleTime uint16) *sqlx.DB {
	databaseUrl := ConnectDBUrl(user, password, dbName, host, port, false)

	db, err := sqlx.Open("mysql", databaseUrl)
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
