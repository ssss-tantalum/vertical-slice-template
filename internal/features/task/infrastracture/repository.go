package infrastracture

import (
	"context"
	"database/sql"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
)

type IRepository interface {
	Create(ctx context.Context, task *task.Task) (uint32, error)
	FindAll(ctx context.Context) ([]task.Task, error)
	FindByID(ctx context.Context, task *task.Task) (*task.Task, error)
	Update(ctx context.Context, task *task.Task) error
	Delete(ctx context.Context, task *task.Task) error
}

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) IRepository {
	return &Repository{
		db: db,
	}
}
