package menu_handler

import (
	category_usecase "api/internal/application/category"
	"api/pkg/httphelper"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type addCategoryReq struct {
	ProjectID    uuid.UUID              `json:"project_id" validate:"required,uuid4"`
	Position     uint                   `json:"position" validate:"required"`
	Translations map[string]*AddInputTr `json:"translations" validate:"required"`
}

type AddInputTr struct {
	Name string `json:"name" validate:"required"`
}

// @Tags         Menu
// @Param        body body addCategoryReq true "body data"
// @Success      200
// @Failure      400  {object}  httphelper.HttpError
// @Router       /projects/{projectID}/menu/category [post]
func (h *handler) addCategory(c echo.Context) error {
	var dto addCategoryReq
	if err := httphelper.BindAndValidate(c, &dto); err != nil {
		return httphelper.HandleError(c, err)
	}

	translations := make(map[string]*category_usecase.AddInputTr)
	for lang, tr := range dto.Translations {
		translations[lang] = &category_usecase.AddInputTr{
			Name: tr.Name,
		}
	}
	input := &category_usecase.AddInput{
		ProjectID:    dto.ProjectID,
		Position:     dto.Position,
		Translations: translations,
	}

	err := h.category.Add(c.Request().Context(), input)
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	return c.JSON(200, "")
}
