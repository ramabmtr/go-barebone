package appctx

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/ramabmtr/go-barebone/app/service/entity"
)

const (
	CtxAuthInfo  = "ctx::auth-info"
	CtxRequestID = "ctx::request-id"
)

func SetEchoRequestID(c echo.Context, rid string) {
	ctx := c.Request().Context()
	ctx = SetRequestID(ctx, rid)
	c.SetRequest(c.Request().Clone(ctx))
	c.Set(CtxRequestID, rid)
}

func SetRequestID(ctx context.Context, rid string) context.Context {
	return context.WithValue(ctx, CtxRequestID, rid)
}

func GetRequestID(ctx context.Context) string {
	rid, _ := ctx.Value(CtxRequestID).(string)
	return rid
}

func SetEchoAuthInfo(c echo.Context, claims *entity.JWTCustomClaims) {
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, CtxAuthInfo, claims)
	c.SetRequest(c.Request().Clone(ctx))
	c.Set(CtxAuthInfo, claims)
}

func GetAuthInfo(ctx context.Context) *entity.JWTCustomClaims {
	claims, _ := ctx.Value(CtxAuthInfo).(*entity.JWTCustomClaims)
	return claims
}
