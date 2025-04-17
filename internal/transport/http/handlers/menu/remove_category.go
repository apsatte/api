package menu_handler

import (
	category_usecase "api/internal/application/category"
	"api/pkg/httphelper"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Tags         Menu
// @Success      200
// @Failure      400  {object}  httphelper.HttpError
// @Router       /projects/{projectID}/menu/category/{categoryID} [delete]
func (h *handler) removeCategory(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("projectID"))
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	categoryID, err := uuid.Parse(c.Param("categoryID"))
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	input := &category_usecase.RemoveInput{
		ProjectID:  projectID,
		CategoryID: categoryID,
	}
	err = h.category.Remove(c.Request().Context(), input)
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	return c.JSON(200, "")
}
