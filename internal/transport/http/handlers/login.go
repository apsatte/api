package account_handler

import (
	account_usecase "api/internal/application/account"
	"api/pkg/httphelper"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Tags         Account
// @Param        customer body loginReq true "login input data"
// @Success      200  {object}  account_usecase.AuthOutput
// @Failure      400  {object}  httphelper.HttpError
// @Router       /account/login [post]
func (h *handler) Login(c echo.Context) error {
	var dto loginReq

	if err := httphelper.BindAndValidate(c, &dto); err != nil {
		return httphelper.HandleError(c, err)
	}

	input := &account_usecase.LoginInput{
		Email:    dto.Email,
		Password: dto.Password,
	}
	output, accessToken, err := h.account.Login(c.Request().Context(), input)
	if err != nil {
		return httphelper.HandleError(c, err)
	}

	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	return c.JSON(200, output)
}
