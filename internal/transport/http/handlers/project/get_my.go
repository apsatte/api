package projects_handler

import (
	"api/pkg/httphelper"

	"github.com/labstack/echo/v4"
)

// @Security     access_token
// @Tags         Projects
// @Success      200  {object}  []project_usecase.ProjectOutput
// @Failure      400  {object}  httphelper.HttpError
// @Router       /projects [get]
func (h *handler) getMy(c echo.Context) error {
	projects, err := h.project.GetByCustomerID(c.Request().Context())
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	return c.JSON(200, projects)
}
