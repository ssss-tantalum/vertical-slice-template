package service

import (
	"context"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task/infrastracture"
)

type IService interface {
	Create(ctx context.Context, name string) (uint32, error)
	FindAll(ctx context.Context) ([]task.Task, error)
	FindByID(ctx context.Context, id uint32) (*task.Task, error)
	Update(ctx context.Context, id uint32, name string, completed bool) error
	Delete(ctx context.Context, id uint32) error
}

type Service struct {
	repo infrastracture.IRepository
}

func New(repo infrastracture.IRepository) IService {
	return &Service{
		repo: repo,
	}
}
