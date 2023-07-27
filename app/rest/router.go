package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ramabmtr/go-barebone/app/config"
	"github.com/ramabmtr/go-barebone/app/errors"
	"github.com/ramabmtr/go-barebone/app/rest/handler"
	"github.com/ramabmtr/go-barebone/app/service/entity"
	"github.com/ramabmtr/go-barebone/app/util/appctx"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
)

func restricted() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entity.JWTCustomClaims)
		},
		SuccessHandler: func(c echo.Context) {
			user, ok := c.Get("user").(*jwt.Token)
			if !ok {
				c.Error(errors.ErrUnauthorized)
				return
			}
			claims, ok := user.Claims.(*entity.JWTCustomClaims)
			if !ok {
				c.Error(errors.ErrUnauthorized)
				return
			}
			appctx.SetAuthInfo(c, claims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, entity.MessageResponse{Message: err.Error()})
		},
		SigningKey: []byte(config.Conf.App.JWT.Secret),
	})
}

func RegisterRouter(e *echo.Echo, h *handler.Handler) {
	e.Any("/ping", h.Ping.Ping)
	e.Any("/ping-restrict", h.Ping.PingRestrict, restricted())

	// docs related
	if config.Conf.FeatureFlag.EnableDocs {
		e.GET("/docs", handler.ServeDoc)
		e.GET("/docs/swagger.yaml", handler.DocsSpec)
	}

	e.POST("/register", h.User.Register)
	e.POST("/login", h.User.Login)
}
