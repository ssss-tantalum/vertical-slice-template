package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ssss-tantalum/vertical-slice-template/internal/infrastracture/router"
	"go.uber.org/fx"
)

type Options struct {
	Port int    `default:"8000"`
	Env  string `default:"dev"`
}

func New(lc fx.Lifecycle, opts *Options) *echo.Echo {
	e := router.New()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", opts.Port),
		Handler: e,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return e
}
