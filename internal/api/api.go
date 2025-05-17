package api

import (
	"context"
	pkgErrors "errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ramabmtr/go-barebone/internal/config"
	"github.com/ramabmtr/go-barebone/internal/errors"
	"github.com/ramabmtr/go-barebone/internal/lib/appctx"
	"github.com/ramabmtr/go-barebone/internal/lib/response"
	"github.com/ramabmtr/go-barebone/internal/service"
	"github.com/rs/zerolog/log"
)

type Rest struct {
	e   *echo.Echo
	svc *service.Service
}

func New(svc *service.Service) *Rest {
	return &Rest{
		e:   echo.New(),
		svc: svc,
	}
}

func skipperByURLPath(paths ...string) func(c echo.Context) bool {
	return func(c echo.Context) bool {
		for _, path := range paths {
			if strings.Contains(c.Request().URL.Path, path) {
				return true
			}
		}
		return false
	}
}

func (r *Rest) Run() {
	r.e.HideBanner = true
	r.e.HidePort = true
	r.e.Validator = config.NewValidator()

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, response.Message("route not found"))
	}

	echo.MethodNotAllowedHandler = func(c echo.Context) error {
		return c.JSON(http.StatusMethodNotAllowed, response.Message("method not allowed"))
	}

	r.e.HTTPErrorHandler = func(err error, c echo.Context) {
		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				err = c.NoContent(errors.ToHTTPCode(err))
			} else {
				err = c.JSON(errors.ToHTTPCode(err), response.Error(err))
			}
		}
	}

	r.e.Pre(middleware.RemoveTrailingSlash())
	r.e.Use(middleware.CORS())
	r.e.Use(middleware.Gzip())
	r.e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		RequestIDHandler: func(e echo.Context, rid string) {
			appctx.SetEchoRequestID(e, rid)
		},
	}))
	r.e.Use(middleware.Recover())
	r.e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper:   skipperByURLPath("docs"),
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().Ctx(c.Request().Context()).
				Str("method", v.Method).
				Str("uri", v.URI).
				Int("status", v.Status).
				Msg("incoming request")
			return nil
		},
	}))

	RegisterRouter(r.e, r.svc)

	go func() {
		log.Info().Str("port", config.GetEnv().Server.Port).Msg("api server started")
		err := r.e.Start(fmt.Sprintf(":%s", config.GetEnv().Server.Port))
		if err != nil && !pkgErrors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("error starting api server")
		}
	}()
}

func (r *Rest) Stop(ctx context.Context) {
	defer log.Info().Msg("api server stopped")
	err := r.e.Shutdown(ctx)
	if err != nil {
		log.Error().Err(err).Msg("error stopping api server")
	}
}
