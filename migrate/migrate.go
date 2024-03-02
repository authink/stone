package migrate

import (
	"fmt"

	"github.com/authink/inkstone/db"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func createSourceUrl(dbMigrateFileSource string) string {
	return fmt.Sprintf("file://%s", dbMigrateFileSource)
}

func Schema(direction, dbMigrateFileSource, dbUser, dbPasswd, dbName, dbHost string, dbPort uint16) {
	if direction != "up" && direction != "down" {
		panic(fmt.Errorf("migrate: unkwon direction %s", direction))
	}

	sourceUrl := createSourceUrl(dbMigrateFileSource)
	databaseUrl := db.MakeUrl(
		dbUser,
		dbPasswd,
		dbName,
		dbHost,
		dbPort,
		true,
	)

	m, err := migrate.New(
		sourceUrl,
		databaseUrl,
	)

	if err != nil {
		panic(err)
	}

	switch direction {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	}

	if err != nil {
		panic(err)
	}
}
