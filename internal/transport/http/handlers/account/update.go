package account_handler

import (
	account_usecase "api/internal/application/account"
	"api/pkg/httphelper"

	"github.com/labstack/echo/v4"
)

// @Tags         Account
// @Param        customer body updateReq true "update input data"
// @Success      200
// @Failure      400  {object}  httphelper.HttpError
// @Router       /account [patch]
func (h *handler) update(c echo.Context) error {
	var dto updateReq

	if err := httphelper.BindAndValidate(c, &dto); err != nil {
		return httphelper.HandleError(c, err)
	}

	input := &account_usecase.UpdateInput{
		Name: dto.Name,
	}
	err := h.account.Update(c.Request().Context(), input)
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	return c.JSON(200, "")
}
