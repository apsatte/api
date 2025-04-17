package menu_handler

import (
	category_usecase "api/internal/application/category"
	"api/pkg/httphelper"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type updateCategoryReq struct {
	Translations map[string]*updateCategoryTr `json:"translations" validate:"required"`
}

type updateCategoryTr struct {
	Name string `json:"name" validate:"required"`
}

// @Tags         Menu
// @Param        body body updateCategoryReq true "body data"
// @Success      200
// @Failure      400  {object}  httphelper.HttpError
// @Router       /projects/{projectID}/menu/category/{categoryID} [patch]
func (h *handler) updateCategory(c echo.Context) error {
	var dto updateCategoryReq
	if err := httphelper.BindAndValidate(c, &dto); err != nil {
		return httphelper.HandleError(c, err)
	}

	categoryID, err := uuid.Parse(c.Param("categoryID"))
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	translations := make(map[string]*category_usecase.UpdateInputTr)
	for lang, tr := range dto.Translations {
		translations[lang] = &category_usecase.UpdateInputTr{
			Name: tr.Name,
		}
	}
	input := &category_usecase.UpdateInput{
		CategoryID:   categoryID,
		Translations: translations,
	}

	err = h.category.Update(c.Request().Context(), input)
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	return c.JSON(200, "")
}
