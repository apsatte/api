package account_handler

import (
	"api/pkg/httphelper"

	"github.com/labstack/echo/v4"
)

// @Security     access_token
// @Tags         Account
// @Success      200  {object}  account_usecase.AuthOutput
// @Failure      400  {object}  httphelper.HttpError
// @Router       /account [get]
func (h *handler) getMe(c echo.Context) error {
	output, err := h.account.GetMe(c.Request().Context())
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	return c.JSON(200, output)
}
