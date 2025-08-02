package handler

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/labstack/echo/v4"
	"github.com/ssss-tantalum/vertical-slice-template/internal/features/task/service"
)

type Handler struct {
	service service.IService
}

func New(service service.IService) *Handler {
	return &Handler{
		service: service,
	}
}

func Routes(e *echo.Echo, api huma.API, h *Handler) {
	huma.Register(api, huma.Operation{
		OperationID:   "create-task",
		Method:        http.MethodPost,
		Path:          "/tasks",
		Summary:       "Create a task",
		Tags:          []string{"Tasks"},
		DefaultStatus: http.StatusCreated,
	}, h.Create)

	huma.Register(api, huma.Operation{
		OperationID: "get-all-task",
		Method:      http.MethodGet,
		Path:        "/tasks",
		Summary:     "Get all task",
		Tags:        []string{"Tasks"},
	}, h.List)

	huma.Register(api, huma.Operation{
		OperationID: "get-task",
		Method:      http.MethodGet,
		Path:        "/tasks/{id}",
		Summary:     "Get a task",
		Tags:        []string{"Tasks"},
	}, h.Show)

	huma.Register(api, huma.Operation{
		OperationID:   "update-task",
		Method:        http.MethodPut,
		Path:          "/tasks/{id}",
		Summary:       "Update a task",
		Tags:          []string{"Tasks"},
		DefaultStatus: http.StatusNoContent,
	}, h.Update)

	huma.Register(api, huma.Operation{
		OperationID:   "delete-task",
		Method:        http.MethodDelete,
		Path:          "/tasks/{id}",
		Summary:       "Delete a task",
		Tags:          []string{"Tasks"},
		DefaultStatus: http.StatusNoContent,
	}, h.Delete)
}
