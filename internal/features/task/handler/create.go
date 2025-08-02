package handler

import (
	"context"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
)

type CreateTaskInput struct {
	Body struct {
		Name string `json:"name"`
	}
}

type CraeteTaskOutput struct {
	Body *task.Task
}

func (h Handler) Create(ctx context.Context, i *CreateTaskInput) (*CraeteTaskOutput, error) {
	taskID, err := h.service.Create(ctx, i.Body.Name)
	if err != nil {
		return nil, err
	}

	task, err := h.service.FindByID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	resp := &CraeteTaskOutput{}
	resp.Body.ID = task.ID
	resp.Body.Name = task.Name
	resp.Body.Completed = task.Completed

	return resp, nil
}
