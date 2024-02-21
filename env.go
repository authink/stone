package inkstone

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
	BasePath             string
}

func LoadEnv() *Env {
	appENV := getAppENV()
	appCWD := getAppCWD()
	appName := "ink"
	secretKey := "your-secret-key"
	accessTokenDuration := uint16(7200)
	refreshTokenDuration := uint16(7 * 24)
	host := "localhost"
	port := uint16(8080)
	shutdownTimeout := uint16(5)

	GetEnvString("APP_NAME", &appName)
	GetEnvString("SECRET_KEY", &secretKey)
	GetEnvUint16("ACCESS_TOKEN_DURATION", &accessTokenDuration)
	GetEnvUint16("REFRESH_TOKEN_DURATION", &refreshTokenDuration)

	GetEnvString("HOST", &host)
	GetEnvUint16("PORT", &port)
	GetEnvUint16("SHUTDOWN_TIMEOUT", &shutdownTimeout)

	dbHost := "localhost"
	dbPort := uint16(3306)
	dbUser := "root"
	dbPasswd := ""
	dbName := "ink"

	GetEnvString("DB_HOST", &dbHost)
	GetEnvUint16("DB_PORT", &dbPort)
	GetEnvString("DB_USER", &dbUser)
	GetEnvString("DB_PASSWORD", &dbPasswd)
	GetEnvString("DB_NAME", &dbName)

	dbMaxOpenConns := uint16(20)
	dbMaxIdleConns := uint16(10)
	dbConnMaxLifeTime := uint16(3600)
	dbConnMaxIdleTime := uint16(300)
	GetEnvUint16("DB_MAX_OPEN_CONNS", &dbMaxOpenConns)
	GetEnvUint16("DB_MAX_IDLE_CONNS", &dbMaxIdleConns)
	GetEnvUint16("DB_CONN_MAX_LIFE_TIME", &dbConnMaxLifeTime)
	GetEnvUint16("DB_CONN_MAX_IDLE_TIME", &dbConnMaxIdleTime)

	dbMigrateFileSource := "../ink.schema/migrations"
	GetEnvString("DB_MIGRATE_FILE_SOURCE", &dbMigrateFileSource)

	basePath := "api/v1"
	GetEnvString("BASE_PATH", &basePath)

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
		basePath,
	}
}

func AssertEnvDev(feature string) {
	if getAppENV() != DEVELOPMENT {
		panic(fmt.Sprintf("[%s] assert development env failed", feature))
	}
}

func getAppENV() string {
	appENV := DEVELOPMENT
	GetEnvString("APP_ENV", &appENV)

	if !(appENV == DEVELOPMENT || appENV == TEST || appENV == PRODUCTION) {
		panic(fmt.Sprintf("Invalid APP_ENV %s", appENV))
	}
	return appENV
}

func getAppCWD() (appCWD string) {
	GetEnvString("APP_CWD", &appCWD)
	return
}

func GetEnvUint16(key string, value *uint16) {
	if v := os.Getenv(key); len(v) > 0 {
		if _, err := fmt.Sscanf(v, "%d", value); err != nil {
			panic(err)
		}
	}
}

func GetEnvString(key string, value *string) {
	if v := os.Getenv(key); len(v) > 0 {
		*value = v
	}
}
