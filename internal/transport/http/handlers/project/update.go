package projects_handler

import (
	project_usecase "api/internal/application/project"
	"api/pkg/httphelper"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Tags         Projects
// @Param        project body updateReq true "update project body"
// @Param        projectID path string true "Project ID"
// @Success      200
// @Failure      400  {object}  httphelper.HttpError
// @Router       /projects/{projectID} [patch]
func (h *handler) update(c echo.Context) error {
	var dto updateReq
	if err := httphelper.BindAndValidate(c, &dto); err != nil {
		return httphelper.HandleError(c, err)
	}

	projectID, err := uuid.Parse(c.Param("projectID"))
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	translations := make(map[string]*project_usecase.UpdateInputTranslation)
	for lang, tr := range dto.Translations {
		translations[lang] = &project_usecase.UpdateInputTranslation{
			Description: tr.Description,
		}
	}

	input := &project_usecase.UpdateInput{
		ProjectID:    projectID,
		Name:         dto.Name,
		ServiceFee:   dto.ServiceFee,
		Languages:    dto.Languages,
		Translations: translations,
	}

	err = h.project.Update(c.Request().Context(), input)
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	return c.JSON(200, "")
}
