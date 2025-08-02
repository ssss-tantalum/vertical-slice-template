package infrastracture

import (
	"context"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
)

func (r Repository) Create(ctx context.Context, task *task.Task) (uint32, error) {
	q := `
		INSERT INTO tasks
			(name)
		VALUES
			(?)
	`

	row, err := r.db.ExecContext(ctx, q, task.Name)
	if err != nil {
		return 0, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}
