package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ramabmtr/go-barebone/app/config"
	"github.com/ramabmtr/go-barebone/app/errors"
	"github.com/ramabmtr/go-barebone/app/rest/handler"
	"github.com/ramabmtr/go-barebone/app/service/entity"
	"github.com/ramabmtr/go-barebone/app/util/appctx"
)

type Rest struct {
	e *echo.Echo
	h *handler.Handler
}

func New(h *handler.Handler) *Rest {
	return &Rest{
		e: echo.New(),
		h: h,
	}
}

func SkipperByURLPath(paths ...string) func(c echo.Context) bool {
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
	r.e.Validator = config.GetValidator()

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, entity.MessageResponse{Message: "route not found"})
	}

	echo.MethodNotAllowedHandler = func(c echo.Context) error {
		return c.JSON(http.StatusMethodNotAllowed, entity.MessageResponse{Message: "method not allowed"})
	}

	r.e.HTTPErrorHandler = func(err error, c echo.Context) {
		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				err = c.NoContent(errors.ErrorToHTTPCode(err))
			} else {
				err = c.JSON(errors.ErrorToHTTPCode(err), entity.MessageResponse{Message: err.Error()})
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
	r.e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisableErrorHandler: true,
	}))
	r.e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper:   SkipperByURLPath("docs"),
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Println(fmt.Sprintf("[REQUEST] [%s] %v, status: %v", v.Method, v.URI, v.Status))
			return nil
		},
	}))

	r.registerRouter()

	go func() {
		err := r.e.Start(fmt.Sprintf(":%s", config.Conf.App.Port))
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()
}

func (r *Rest) Stop(ctx context.Context) {
	defer log.Println("rest stopped")
	err := r.e.Shutdown(ctx)
	if err != nil {
		log.Printf("error stopping rest. %s", err.Error())
	}
}
