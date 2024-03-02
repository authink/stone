package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if appCWD := getAppCWD(); appCWD != "" {
		if err := os.Chdir(appCWD); err != nil {
			panic(err)
		}
	}

	appENV := getAppENV()

	files := []string{
		fmt.Sprintf(".env.%s.local", appENV),
		".env.local",
		fmt.Sprintf(".env.%s", appENV),
		".env",
	}

	var existFiles []string
	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			existFiles = append(existFiles, file)
		}
	}

	if len(existFiles) > 0 {
		if err := godotenv.Load(existFiles...); err != nil {
			panic(err)
		}
	}
}

const (
	DEVELOPMENT string = "dev"
	TEST        string = "test"
	PRODUCTION  string = "prod"
)

const (
	APP_ENV                string = "APP_ENV"
	APP_CWD                string = "APP_CWD"
	APP_NAME               string = "APP_NAME"
	SECRET_KEY             string = "SECRET_KEY"
	ACCESS_TOKEN_DURATION  string = "ACCESS_TOKEN_DURATION"
	REFRESH_TOKEN_DURATION string = "REFRESH_TOKEN_DURATION"
	HOST                   string = "HOST"
	PORT                   string = "PORT"
	SHUTDOWN_TIMEOUT       string = "SHUTDOWN_TIMEOUT"
	DB_HOST                string = "DB_HOST"
	DB_PORT                string = "DB_PORT"
	DB_USER                string = "DB_USER"
	DB_PASSWORD            string = "DB_PASSWORD"
	DB_NAME                string = "DB_NAME"
	DB_MAX_OPEN_CONNS      string = "DB_MAX_OPEN_CONNS"
	DB_MAX_IDLE_CONNS      string = "DB_MAX_IDLE_CONNS"
	DB_CONN_MAX_LIFE_TIME  string = "DB_CONN_MAX_LIFE_TIME"
	DB_CONN_MAX_IDLE_TIME  string = "DB_CONN_MAX_IDLE_TIME"
	DB_MIGRATE_FILE_SOURCE string = "DB_MIGRATE_FILE_SOURCE"
	DB_LOG_MODE            string = "DB_LOG_MODE"
	BASE_PATH              string = "BASE_PATH"
)

type Env struct {
	AppENV               string
	AppCWD               string
	AppName              string
	SecretKey            string
	AccessTokenDuration  uint16
	RefreshTokenDuration uint16
	Host                 string
	Port                 uint16
	ShutdownTimeout      uint16
	DbHost               string
	DbPort               uint16
	DbUser               string
	DbPasswd             string
	DbName               string
	DbMaxOpenConns       uint16
	DbMaxIdleConns       uint16
	DbConnMaxLifeTime    uint16
	DbConnMaxIdleTime    uint16
	DbMigrateFileSource  string
	DbLogMode            bool
	BasePath             string
}

func Load() *Env {
	appENV := getAppENV()
	appCWD := getAppCWD()
	appName := "ink"
	secretKey := ""
	accessTokenDuration := uint16(7200)
	refreshTokenDuration := uint16(7 * 24)
	host := "localhost"
	port := uint16(8080)
	shutdownTimeout := uint16(5)

	GetString(APP_NAME, &appName)
	GetString(SECRET_KEY, &secretKey)
	if secretKey == "" {
		panic(fmt.Sprintf("env %s can't be empty", SECRET_KEY))
	}

	GetUint16(ACCESS_TOKEN_DURATION, &accessTokenDuration)
	GetUint16(REFRESH_TOKEN_DURATION, &refreshTokenDuration)

	GetString(HOST, &host)
	GetUint16(PORT, &port)
	GetUint16(SHUTDOWN_TIMEOUT, &shutdownTimeout)

	dbHost := "localhost"
	dbPort := uint16(3306)
	dbUser := "root"
	dbPasswd := ""
	dbName := "ink"

	GetString(DB_HOST, &dbHost)
	GetUint16(DB_PORT, &dbPort)
	GetString(DB_USER, &dbUser)
	GetString(DB_PASSWORD, &dbPasswd)
	GetString(DB_NAME, &dbName)

	dbMaxOpenConns := uint16(20)
	dbMaxIdleConns := uint16(10)
	dbConnMaxLifeTime := uint16(3600)
	dbConnMaxIdleTime := uint16(300)
	GetUint16(DB_MAX_OPEN_CONNS, &dbMaxOpenConns)
	GetUint16(DB_MAX_IDLE_CONNS, &dbMaxIdleConns)
	GetUint16(DB_CONN_MAX_LIFE_TIME, &dbConnMaxLifeTime)
	GetUint16(DB_CONN_MAX_IDLE_TIME, &dbConnMaxIdleTime)

	dbMigrateFileSource := "../ink.schema/migrations"
	GetString(DB_MIGRATE_FILE_SOURCE, &dbMigrateFileSource)

	dbLogMode := appENV == DEVELOPMENT
	GetBool(DB_LOG_MODE, &dbLogMode)

	basePath := "api/v1"
	GetString(BASE_PATH, &basePath)

	return &Env{
		appENV,
		appCWD,
		appName,
		secretKey,
		accessTokenDuration,
		refreshTokenDuration,
		host,
		port,
		shutdownTimeout,
		dbHost,
		dbPort,
		dbUser,
		dbPasswd,
		dbName,
		dbMaxOpenConns,
		dbMaxIdleConns,
		dbConnMaxLifeTime,
		dbConnMaxIdleTime,
		dbMigrateFileSource,
		dbLogMode,
		basePath,
	}
}

func AssertDev(feature string) {
	if getAppENV() != DEVELOPMENT {
		panic(fmt.Sprintf("[%s] assert development env failed", feature))
	}
}

func GetBool(key string, value *bool) {
	if v := os.Getenv(key); len(v) > 0 {
		if _, err := fmt.Sscanf(v, "%t", value); err != nil {
			panic(err)
		}
	}
}

func GetUint16(key string, value *uint16) {
	if v := os.Getenv(key); len(v) > 0 {
		if _, err := fmt.Sscanf(v, "%d", value); err != nil {
			panic(err)
		}
	}
}

func GetString(key string, value *string) {
	if v := os.Getenv(key); len(v) > 0 {
		*value = v
	}
}

func getAppENV() string {
	appENV := DEVELOPMENT
	GetString(APP_ENV, &appENV)

	if !(appENV == DEVELOPMENT || appENV == TEST || appENV == PRODUCTION) {
		panic(fmt.Sprintf("Invalid %s %s", APP_ENV, appENV))
	}
	return appENV
}

func getAppCWD() (appCWD string) {
	GetString(APP_CWD, &appCWD)
	return
}
