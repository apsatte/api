package projects_handler

import (
	"api/pkg/httphelper"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Tags         Projects
// @Param        projectID path string true "Project ID"
// @Success      200
// @Failure      400  {object}  httphelper.HttpError
// @Router       /projects/{projectID} [delete]
func (h *handler) remove(c echo.Context) error {
	projectID, err := uuid.Parse(c.Param("projectID"))
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	err = h.project.Remove(c.Request().Context(), projectID)
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	return c.JSON(200, "")
}
