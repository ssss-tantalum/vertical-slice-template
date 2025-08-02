package handler

import (
	"context"
)

type UpdateTaskInput struct {
	ID   uint32 `path:"id"`
	Body struct {
		Name      string `json:"name"`
		Completed bool   `json:"completed"`
	}
}

func (h Handler) Update(ctx context.Context, i *UpdateTaskInput) (*struct{}, error) {
	err := h.service.Update(ctx, i.ID, i.Body.Name, i.Body.Completed)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
