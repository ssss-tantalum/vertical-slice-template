package handler

import (
	"context"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
)

type ShowTaskInput struct {
	ID uint32 `path:"id"`
}

type ShowTaskOutput struct {
	Body *task.Task
}

func (h Handler) Show(ctx context.Context, i *ShowTaskInput) (*ShowTaskOutput, error) {
	task, err := h.service.FindByID(ctx, i.ID)
	if err != nil {
		return nil, err
	}

	resp := &ShowTaskOutput{}
	resp.Body.ID = task.ID
	resp.Body.Name = task.Name
	resp.Body.Completed = task.Completed

	return resp, nil
}
