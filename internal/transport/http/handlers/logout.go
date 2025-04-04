package account_handler

import (
	"api/pkg/httphelper"

	"github.com/labstack/echo/v4"
)

// @Security     access_token
// @Tags         Account
// @Success      200
// @Router       /account/logout [post]
func (h *handler) logout(c echo.Context) error {
	err := h.account.Logout(c.Request().Context())
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	return c.JSON(200, "")
}
