package service

import (
	"context"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
)

func (s Service) Create(ctx context.Context, name string) (uint32, error) {
	t := &task.Task{
		Name: name,
	}

	id, err := s.repo.Create(ctx, t)
	if err != nil {
		return 0, err
	}

	return id, nil
}
