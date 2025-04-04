package httphelper

import (
	"api/internal/domain"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Middleware interface {
	Authenticate(next echo.HandlerFunc) echo.HandlerFunc
}

type Handler interface {
	Register(g *echo.Group, m Middleware)
}

var v = validator.New()

func BindAndValidate(c echo.Context, dto any) error {
	if err := c.Bind(&dto); err != nil {
		return errors.New("INVALID_INPUT_DATA")
	}

	if err := v.Struct(dto); err != nil {
		return errors.New("INVALID_INPUT_DATA")
	}

	return nil
}

func HandleError(c echo.Context, err error) error {
	if errors.Is(err, domain.ErrDatabase) {
		return c.JSON(500, HttpError{Code: "INTERNAL_SERVER_ERROR"})
	}

	if errors.Is(err, domain.ErrNotFound) {
		return c.JSON(404, HttpError{Code: err.Error()})
	}

	if errors.Is(err, domain.ErrUnauthorized) {
		return c.JSON(401, HttpError{Code: err.Error()})
	}

	return c.JSON(400, HttpError{Code: err.Error()})
}
