package test

import (
	"testing"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task/mock"
	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task/service"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockIRepository(ctrl)
	mockRepo.EXPECT().Delete(t.Context(), &task.Task{ID: 1}).Return(nil)

	service := service.New(mockRepo)

	err := service.Delete(t.Context(), 1)
	assert.NoError(t, err)
}
