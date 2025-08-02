package api

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo) huma.API {
	api := humaecho.New(e, huma.DefaultConfig("Task API", "0.0.1"))
	grp := huma.NewGroup(api, "/api")

	return grp
}
