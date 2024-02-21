package inkstone

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type SeedFunc func(*AppContext)

func createSourceUrl(app *AppContext) string {
	return fmt.Sprintf("file://%s", app.Env.DbMigrateFileSource)
}

func migrateSchema(app *AppContext, direction string) {
	if direction != "up" && direction != "down" {
		panic(fmt.Errorf("migrate: unkwon direction %s", direction))
	}

	sourceUrl := createSourceUrl(app)
	databaseUrl := ConnectDBUrl(
		app.Env.DbUser,
		app.Env.DbPasswd,
		app.Env.DbName,
		app.Env.DbHost,
		app.Env.DbPort,
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