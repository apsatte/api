package projects_handler

import (
	project_usecase "api/internal/application/project"
	"api/pkg/httphelper"
	"encoding/base64"

	"github.com/labstack/echo/v4"
)

// @Tags         Projects
// @Param        project body addReq true "add project data"
// @Success      200
// @Failure      400  {object}  httphelper.HttpError
// @Router       /projects [post]
func (h *handler) add(c echo.Context) error {
	var dto addReq

	if err := httphelper.BindAndValidate(c, &dto); err != nil {
		return httphelper.HandleError(c, err)
	}

	logoBytes, err := base64.StdEncoding.DecodeString(dto.Logo)
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	backgroundBytes, err := base64.StdEncoding.DecodeString(dto.Background)
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	input := project_usecase.NewAddProjectInput(
		dto.Name, dto.Description,
		logoBytes, backgroundBytes,
		dto.ServiceFee, dto.Languages,
	)
	err = h.project.Add(c.Request().Context(), input)
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	return c.JSON(200, "")
}
