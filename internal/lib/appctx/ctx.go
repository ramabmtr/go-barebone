package appctx

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/ramabmtr/go-barebone/internal/entity"
)

const (
	CtxUserInfo  = "ctx::user-info"
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

func SetEchoUserInfo(c echo.Context, userCtx *entity.UserCtx) {
	ctx := c.Request().Context()
	ctx = SetUserInfo(ctx, userCtx)
	c.SetRequest(c.Request().Clone(ctx))
	c.Set(CtxUserInfo, userCtx)
}

func SetUserInfo(ctx context.Context, userCtx *entity.UserCtx) context.Context {
	return context.WithValue(ctx, CtxUserInfo, userCtx)
}

func GetUserInfo(ctx context.Context) *entity.UserCtx {
	userCtx, _ := ctx.Value(CtxUserInfo).(*entity.UserCtx)
	return userCtx
}
