package service

import (
	"context"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
)

func (s Service) Update(ctx context.Context, id uint32, name string, completed bool) error {
	t := &task.Task{
		ID:        id,
		Name:      name,
		Completed: completed,
	}

	err := s.repo.Update(ctx, t)
	if err != nil {
		return err
	}

	return nil
}
