package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ramabmtr/go-barebone/app/errors"
	"github.com/ramabmtr/go-barebone/app/service/entity"
	"github.com/ramabmtr/go-barebone/app/util/appctx"
)

type ping struct{}

func newPingHandler() *ping {
	return &ping{}
}

// Ping :
// @description Return simple response to indicates the app is up and running
// @tags HealthCheck
// @accept json
// @produce json
// @success 200 {object} entity.MessageResponse
// @router /ping [get]
func (h *ping) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, entity.MessageResponse{Message: "pong"})
}

// PingRestrict :
// @description Another Ping to demonstrate auth
// @tags HealthCheck
// @accept json
// @produce json
// @security BearerAuth
// @success 200 {object} entity.MessageResponse
// @router /ping-restrict [get]
func (h *ping) PingRestrict(c echo.Context) error {
	ctx := c.Request().Context()
	authInfo, err := appctx.GetAuthInfo(ctx)
	if err != nil {
		return c.JSON(errors.ErrorToHTTPCode(err), entity.MessageResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, entity.MessageResponse{Message: authInfo.Username})
}
