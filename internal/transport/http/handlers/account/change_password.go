package account_handler

import (
	account_usecase "api/internal/application/account"
	"api/pkg/httphelper"

	"github.com/labstack/echo/v4"
)

// @Tags         Account
// @Param        customer body changePasswordReq true "change password data"
// @Success      200  {object}  account_usecase.AuthOutput
// @Failure      400  {object}  httphelper.HttpError
// @Router       /account/password [patch]
func (h *handler) changePassword(c echo.Context) error {
	var dto changePasswordReq

	if err := httphelper.BindAndValidate(c, &dto); err != nil {
		return httphelper.HandleError(c, err)
	}

	input := &account_usecase.UpdatePasswordInput{
		OldPassword:  dto.OldPassword,
		NewPassword:  dto.NewPassword,
		ClearSession: dto.ClearSession,
	}
	err := h.account.UpdatePassword(c.Request().Context(), input)
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	return nil
}
