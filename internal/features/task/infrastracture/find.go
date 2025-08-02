package infrastracture

import (
	"context"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
)

func (r Repository) FindAll(ctx context.Context) ([]task.Task, error) {
	q := `
		SELECT
			id,
			name,
			completed
		FROM
			tasks
	`

	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var task task.Task
		err := rows.Scan(&task.ID, &task.Name, &task.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r Repository) FindByID(ctx context.Context, t *task.Task) (*task.Task, error) {
	q := `
		SELECT
			id,
			name,
			completed
		FROM
			tasks
		WHERE
			id = ?
	`

	row, err := r.db.QueryContext(ctx, q, t.ID)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var task task.Task
	for row.Next() {
		err := row.Scan(&task.ID, &task.Name, &task.Completed)
		if err != nil {
			return nil, err
		}
	}

	return &task, nil
}
