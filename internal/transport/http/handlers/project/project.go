package projects_handler

import (
	project_usecase "api/internal/application/project"
	"api/pkg/httphelper"
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProjectUseCase interface {
	Add(c context.Context, d *project_usecase.AddProjectInput) error
	GetByCustomerID(c context.Context) ([]*project_usecase.ProjectOutput, error)
	Remove(c context.Context, projectID uuid.UUID) error
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
	private.GET("", h.getMy)
	private.POST("", h.add)
	private.DELETE("/:projectID", h.remove)
}
