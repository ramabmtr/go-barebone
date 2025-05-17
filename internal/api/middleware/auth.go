package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/ramabmtr/go-barebone/internal/entity"
	"github.com/ramabmtr/go-barebone/internal/lib/appctx"
)

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get(echo.HeaderAuthorization)
			if token == "" {
				return echo.ErrUnauthorized
			}
			// set dummy user info
			appctx.SetEchoUserInfo(c, &entity.UserCtx{
				ID:       "1",
				Username: "dummy",
				Role:     "USER",
			})
			return next(c)
		}
	}
}
