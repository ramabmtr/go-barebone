package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ramabmtr/go-barebone/internal/lib/response"
	"github.com/ramabmtr/go-barebone/internal/service"
)

type Dummy struct {
	dummySvc *service.Dummy
}

func NewDummyHandler(dummySvc *service.Dummy) *Dummy {
	return &Dummy{
		dummySvc: dummySvc,
	}
}

func (h *Dummy) Dummy(c echo.Context) error {
	ctx := c.Request().Context()
	err := h.dummySvc.Dummy(ctx, "api")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(err))
	}
	return c.JSON(http.StatusOK, response.Message("success"))
}
