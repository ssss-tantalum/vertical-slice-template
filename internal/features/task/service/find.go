package service

import (
	"context"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
)

func (s Service) FindAll(ctx context.Context) ([]task.Task, error) {
	return s.repo.FindAll(ctx)
}

func (s Service) FindByID(ctx context.Context, id uint32) (*task.Task, error) {
	t := &task.Task{
		ID: id,
	}

	task, err := s.repo.FindByID(ctx, t)
	if err != nil {
		return nil, err
	}

	return task, nil
}
