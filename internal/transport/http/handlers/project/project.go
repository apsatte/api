package projects_handler

import (
	project_usecase "api/internal/application/project"
	"api/pkg/httphelper"
	"context"

	"github.com/labstack/echo/v4"
)

type ProjectUseCase interface {
	Add(c context.Context, d *project_usecase.AddProjectInput) error
}

type handler struct {
	project ProjectUseCase
}

func New(project ProjectUseCase) *handler {
	return &handler{
		project: project,
	}
}

func (h *handler) Register(g *echo.Group, m httphelper.Middleware) {
	projects := g.Group("/projects")

	private := projects.Group("", m.Authenticate)
	private.POST("", h.add)
}
