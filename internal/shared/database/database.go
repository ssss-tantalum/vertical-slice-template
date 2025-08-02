package database

import (
	"context"
	"database/sql"

	"github.com/ssss-tantalum/vertical-slice-template/internal/shared/config"
	"go.uber.org/fx"

	_ "github.com/go-sql-driver/mysql"
)

func New(lc fx.Lifecycle, cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	return db, nil
}
