package test

import (
	"testing"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task/mock"
	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task/service"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreate_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockIRepository(ctrl)
	mockRepo.EXPECT().Create(t.Context(), &task.Task{Name: "task1"}).Return(uint32(1), nil)

	service := service.New(mockRepo)

	id, err := service.Create(t.Context(), "task1")
	assert.NoError(t, err)
	if id != 1 {
		t.Errorf("expected task ID to be '1', got: %v", id)
	}
}
