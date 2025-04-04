package account_handler

import (
	account_usecase "api/internal/application/account"
	"api/pkg/httphelper"
	"context"

	"github.com/labstack/echo/v4"
)

type AccountUseCase interface {
	Login(c context.Context, d *account_usecase.LoginInput) (*account_usecase.AuthOutput, string, error)
	Logout(c context.Context) error
	Update(c context.Context, d *account_usecase.UpdateInput) error
	UpdatePassword(c context.Context, d *account_usecase.UpdatePasswordInput) error
	GetMe(c context.Context) (*account_usecase.AuthOutput, error)
}

type handler struct {
	account AccountUseCase
}

func New(account AccountUseCase) httphelper.Handler {
	return &handler{
		account: account,
	}
}

func (h *handler) Register(g *echo.Group, m httphelper.Middleware) {
	account := g.Group("/account")

	account.POST("/login", h.Login)

	{
		private := account.Group("", m.Authenticate)
		private.GET("", h.getMe)
		private.POST("/logout", h.logout)
		private.PATCH("", h.update)
		private.PATCH("/password", h.changePassword)
	}
}
