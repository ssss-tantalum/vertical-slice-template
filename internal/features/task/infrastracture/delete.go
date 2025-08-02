package infrastracture

import (
	"context"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
)

func (r Repository) Delete(ctx context.Context, task *task.Task) error {
	q := `
		DELETE FROM
			tasks
		WHERE
			id = ?
	`

	_, err := r.db.ExecContext(ctx, q, task.ID)
	if err != nil {
		return err
	}

	return nil
}
