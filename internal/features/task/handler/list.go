package handler

import (
	"context"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
)

type ListTaskOutput struct {
	Body struct {
		Tasks []task.Task
	}
}

func (h Handler) List(ctx context.Context, _ *struct{}) (*ListTaskOutput, error) {
	tasks, err := h.service.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	resp := &ListTaskOutput{}
	resp.Body.Tasks = tasks

	return resp, nil
}
