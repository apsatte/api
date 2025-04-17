package menu_handler

import (
	category_usecase "api/internal/application/category"
	dish_usecase "api/internal/application/dish"
	"api/pkg/httphelper"
	"context"

	"github.com/labstack/echo/v4"
)

type CategoryUseCase interface {
	Add(c context.Context, dto *category_usecase.AddInput) error
	Remove(c context.Context, dto *category_usecase.RemoveInput) error
	Update(c context.Context, dto *category_usecase.UpdateInput) error
	UpdatePosition(c context.Context, dto *category_usecase.UpdatePositionInput) error
}

type DishUseCase interface {
	Add(c context.Context, dto *dish_usecase.AddInput) error
}

type handler struct {
	category CategoryUseCase
	dish     DishUseCase
}

func New(
	category CategoryUseCase,
	dish DishUseCase) httphelper.Handler {
	return &handler{
		category: category,
		dish:     dish,
	}
}

func (h *handler) Register(g *echo.Group, m httphelper.Middleware) {
	menu := g.Group("/projects/:projectID/menu", m.Authenticate)

	category := menu.Group("/category")
	category.POST("", h.addCategory)
	category.DELETE("/:categoryID", h.removeCategory)
	category.PATCH("/:categoryID", h.updateCategory)
	category.PATCH("/:categoryID/positions", h.updateCategoryPosition)

	dish := category.Group("/:categoryID")
	dish.POST("", func(c echo.Context) error { return nil })
}
