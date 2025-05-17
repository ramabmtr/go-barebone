package api

import (
	"github.com/labstack/echo/v4"
	"github.com/ramabmtr/go-barebone/internal/api/handler"
	"github.com/ramabmtr/go-barebone/internal/api/middleware"
	"github.com/ramabmtr/go-barebone/internal/service"
)

func RegisterRouter(e *echo.Echo, svc *service.Service) {
	dummyHandler := handler.NewDummyHandler(svc.Dummy)

	apiGroup := e.Group("/api", middleware.Auth())

	apiGroup.GET("/dummy", dummyHandler.Dummy)
}
