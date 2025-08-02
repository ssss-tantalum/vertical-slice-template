package service

import (
	"testing"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestFindByID_Success(t *testing.T) {
	resp := &task.Task{
		ID:        1,
		Name:      "task1",
		Completed: false,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockIRepository(ctrl)
	mockRepo.EXPECT().FindByID(t.Context(), &task.Task{ID: 1}).Return(resp, nil)

	service := New(mockRepo)

	task, err := service.FindByID(t.Context(), 1)
	assert.NoError(t, err)
	if task.ID != 1 {
		t.Errorf("expected task ID to be '1', got: %v", task.ID)
	}
	if task.Name != "task1" {
		t.Errorf("expected task name to be 'task1', got: %v", task.Name)
	}
	if task.Completed != false {
		t.Errorf("expected task completed to be 'false', got: %v", task.Name)
	}
}

func TestFindAll_Success(t *testing.T) {
	resp := []task.Task{
		{ID: 1, Name: "task1", Completed: false},
		{ID: 2, Name: "task2", Completed: true},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockIRepository(ctrl)
	mockRepo.EXPECT().FindAll(t.Context()).Return(resp, nil)

	service := New(mockRepo)

	tasks, err := service.FindAll(t.Context())
	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
}
