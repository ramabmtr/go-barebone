package appctx

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/ramabmtr/go-barebone/app/errors"
	"github.com/ramabmtr/go-barebone/app/service/entity"
)

func SetAuthInfo(c echo.Context, claims *entity.JWTCustomClaims) {
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, entity.CtxAuthInfo, claims)
	c.SetRequest(c.Request().Clone(ctx))
	c.Set(entity.CtxAuthInfo, claims)
}

func GetAuthInfo(ctx context.Context) (*entity.JWTCustomClaims, error) {
	claims, ok := ctx.Value(entity.CtxAuthInfo).(*entity.JWTCustomClaims)
	if !ok || claims == nil {
		return nil, errors.ErrUnauthorized
	}

	return claims, nil
}
