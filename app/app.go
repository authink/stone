package app

import (
	"log"
	"net/http"
	"os"

	"github.com/authink/inkstone/env"
	"github.com/authink/inkstone/migrate"
	"github.com/authink/inkstone/server"
	"github.com/cosmtrek/air/runner"
	"github.com/spf13/cobra"
	"github.com/swaggo/swag"
	"github.com/swaggo/swag/format"
	"github.com/swaggo/swag/gen"
)

type CreateHandlerFunc func() http.Handler

func Run(createHandler CreateHandlerFunc, appCtx *AppContext, opts *Options) {
	var cmd = &cobra.Command{Use: appCtx.AppName}

	var cmdMigrate = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate schema up or down",
		Run: func(cmd *cobra.Command, args []string) {
			direction, err := cmd.Flags().GetString("direction")
			if err != nil {
				panic(err)
			}
			migrate.Schema(
				direction,
				appCtx.DbMigrateFileSource,
				appCtx.DbUser,
				appCtx.DbPasswd,
				appCtx.DbName,
				appCtx.DbHost,
				appCtx.DbPort,
			)
		},
	}

	var cmdSeed = &cobra.Command{
		Use:   "seed",
		Short: "Seed the database",
		Run: func(cmd *cobra.Command, args []string) {
			if opts.Seed != nil {
				opts.Seed(appCtx)
			}
		},
	}

	var cmdSwag = &cobra.Command{
		Use:   "swag",
		Short: "Generate swagger docs",
		Run: func(cmd *cobra.Command, args []string) {
			err := format.New().Build(&format.Config{
				SearchDir: "./src",
				MainFile:  "router/setup.go",
			})
			if err != nil {
				panic(err)
			}

			err = gen.New().Build(&gen.Config{
				SearchDir:   "./src",
				MainAPIFile: "router/setup.go",
				OutputDir:   "./src/docs",

				PropNamingStrategy: swag.CamelCase,
				OutputTypes:        []string{"go", "json", "yaml"},

				ParseDependency: 1,
				ParseDepth:      100,
				ParseGoList:     true,

				OverridesFile:      gen.DefaultOverridesFile,
				LeftTemplateDelim:  "{{",
				RightTemplateDelim: "}}",

				Debugger:         log.New(os.Stdout, "", log.LstdFlags),
				CollectionFormat: swag.TransToValidCollectionFormat("csv"),
			})
			if err != nil {
				panic(err)
			}
		},
	}

	var cmdRun = &cobra.Command{
		Use:   "run",
		Short: "Run server",
		Run: func(cmd *cobra.Command, args []string) {
			hotReload, err := cmd.Flags().GetBool("live-reload")

			if err != nil {
				panic(err)
			}

			if hotReload {
				env.AssertDev("live-reload")

				cfg, err := runner.InitConfig("")
				if err != nil {
					panic(err)
				}

				cfg.Build.Cmd = "go build -o ./tmp/main ./src"
				cfg.Build.ArgsBin = []string{"run"}

				r, err := runner.NewEngineWithConfig(cfg, true)
				if err != nil {
					panic(err)
				}

				r.Run()
			} else {
				server.Start(
					appCtx.Host,
					appCtx.Port,
					appCtx.ShutdownTimeout,
					createHandler(),
				)
			}
		},
	}

	cmdMigrate.Flags().StringP("direction", "d", "up", "Specify migrate direction[up, down]")
	cmdRun.Flags().BoolP("live-reload", "l", false, "Enable live reload")

	cmd.AddCommand(cmdMigrate)
	cmd.AddCommand(cmdSeed)
	cmd.AddCommand(cmdSwag)
	cmd.AddCommand(cmdRun)

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
