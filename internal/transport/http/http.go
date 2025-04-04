package http

import (
	"api/internal/domain/customer"
	"api/pkg/httphelper"
	"context"

	_ "api/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title        Foodx API Documentation
// @version      1.0
// @description  no description
// @BasePath     /api/v1
// @Accept       json
// @Produce      json

type httpServer struct {
	sessionsRepo customer.SessionsRepository
	e            *echo.Echo
}

func New(sessionsRepo customer.SessionsRepository) *httpServer {
	return &httpServer{
		e:            echo.New(),
		sessionsRepo: sessionsRepo,
	}
}

func (h *httpServer) Run(addr string) error {
	return h.e.Start(addr)
}

func (h *httpServer) Shutdown(c context.Context) error {
	return h.e.Shutdown(c)
}

func (h *httpServer) SetupRouter(handlers ...httphelper.Handler) {
	h.e.GET("/swagger/*", echoSwagger.WrapHandler)

	//
	baseUrl := h.e.Group("/api/v1")

	for _, handler := range handlers {
		handler.Register(baseUrl, h)
	}
}
