package test

import (
	"testing"

	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task"
	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task/mock"
	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task/service"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockIRepository(ctrl)
	mockRepo.EXPECT().Update(t.Context(), &task.Task{ID: 1, Name: "task1", Completed: true}).Return(nil)

	service := service.New(mockRepo)

	err := service.Update(t.Context(), 1, "task1", true)
	assert.NoError(t, err)
}
