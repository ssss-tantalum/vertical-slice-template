package handler

import (
	"context"
)

type DeleteTaskInput struct {
	ID uint32 `path:"id"`
}

func (h Handler) Delete(ctx context.Context, i *DeleteTaskInput) (*struct{}, error) {
	err := h.service.Delete(ctx, i.ID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
