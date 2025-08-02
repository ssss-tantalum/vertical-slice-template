package infrastracture

import (
	"context"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
)

func (r Repository) Update(ctx context.Context, task *task.Task) error {
	q := `
		UPDATE
			tasks
		SET
			name = ?,
			completed = ?
		WHERE
			id = ?
	`

	_, err := r.db.ExecContext(ctx, q, task.Name, task.Completed, task.ID)
	if err != nil {
		return err
	}

	return nil
}
