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
			appctx.SetEchoAuthInfo(c, claims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, entity.MessageResponse{Message: err.Error()})
		},
		SigningKey: []byte(config.Conf.App.JWT.Secret),
	})
}

func (r *Rest) registerRouter() {
	r.e.Any("/ping", r.h.Ping.Ping)
	r.e.Any("/ping-restrict", r.h.Ping.PingRestrict, restricted())

	// docs related
	if config.Conf.FeatureFlag.EnableDocs {
		r.e.GET("/docs", handler.ServeDoc)
		r.e.GET("/docs/swagger.yaml", handler.DocsSpec)
	}

	r.e.POST("/register", r.h.User.Register)
	r.e.POST("/login", r.h.User.Login)
}
