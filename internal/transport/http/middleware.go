package http

import (
	"api/internal/domain"
	"api/internal/domain/customer"
	"api/pkg/httphelper"
	"errors"

	"github.com/labstack/echo/v4"
)

func (h *httpServer) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := c.Request().Header.Get("Authorization")

		if accessToken == "" {
			token, err := c.Cookie("access_token")
			if err != nil {
				return httphelper.HandleError(c, domain.ErrUnauthorized)
			}

			accessToken = token.Value
		}

		session, err := h.sessionsRepo.GetOneByAccessToken(c.Request().Context(), accessToken)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				return httphelper.HandleError(c, domain.ErrUnauthorized)
			}
			return httphelper.HandleError(c, err)
		}

		ctx := customer.SetSessionToContext(c.Request().Context(), session)

		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
