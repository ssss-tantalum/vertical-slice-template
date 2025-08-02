package service

import (
	"context"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
)

func (s Service) Delete(ctx context.Context, id uint32) error {
	t := &task.Task{
		ID: id,
	}

	return s.repo.Delete(ctx, t)
}
