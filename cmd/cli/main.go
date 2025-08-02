package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/labstack/echo/v4"
	taskHandler "github.com/ssss-tantalum/vertical-slice-template/internal/features/task/handler"
	taskRepo "github.com/ssss-tantalum/vertical-slice-template/internal/features/task/infrastracture"
	taskService "github.com/ssss-tantalum/vertical-slice-template/internal/features/task/service"
	"github.com/ssss-tantalum/vertical-slice-template/internal/shared/api"
	"github.com/ssss-tantalum/vertical-slice-template/internal/shared/config"
	"github.com/ssss-tantalum/vertical-slice-template/internal/shared/database"
	"github.com/ssss-tantalum/vertical-slice-template/internal/shared/server"
	"go.uber.org/fx"
)

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, opts *server.Options) {
		cfg, err := config.ReadConfig(opts.Env)
		if err != nil {
			log.Fatal(err)
		}

		app := fx.New(
			fx.Supply(cfg),
			fx.Supply(opts),
			fx.Provide(
				// Task
				taskRepo.New,
				taskService.New,
				taskHandler.New,
				// infra
				database.New,
				server.New,
				api.New,
			),
			fx.Invoke(
				taskHandler.Routes,
				func(*sql.DB) {},
				func(*echo.Echo) {},
			),
		)

		hooks.OnStart(func() {
			app.Run()
		})
		hooks.OnStop(func() {
			app.Stop(context.Background())
		})
	})

	cli.Run()
}
