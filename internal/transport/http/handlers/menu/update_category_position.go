package menu_handler

import (
	category_usecase "api/internal/application/category"
	"api/pkg/httphelper"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type updateCategoryPositionReq struct {
	Categories []updateCategoryPositionReqItem `json:"categories" validate:"required"`
}

type updateCategoryPositionReqItem struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Position uint      `json:"position" validate:"required"`
}

func mapToUpdatePositionInput(req updateCategoryPositionReq) *category_usecase.UpdatePositionInput {
	categories := make([]category_usecase.UpdatePositionInputCategory, len(req.Categories))
	for i, item := range req.Categories {
		categories[i] = category_usecase.UpdatePositionInputCategory{
			ID:       item.ID,
			Position: item.Position,
		}
	}
	return &category_usecase.UpdatePositionInput{
		Categories: categories,
	}
}

// @Tags         Menu
// @Param        body body updateCategoryPositionReq true "body data"
// @Success      200
// @Failure      400  {object}  httphelper.HttpError
// @Router       /projects/{projectID}/menu/category/{categoryID}/positions [patch]
func (h *handler) updateCategoryPosition(c echo.Context) error {
	var dto updateCategoryPositionReq

	if err := httphelper.BindAndValidate(c, &dto); err != nil {
		return httphelper.HandleError(c, err)
	}

	input := mapToUpdatePositionInput(dto)
	err := h.category.UpdatePosition(c.Request().Context(), input)
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	return c.JSON(200, "")
}
