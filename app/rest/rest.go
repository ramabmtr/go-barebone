package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ramabmtr/go-barebone/app/config"
	"github.com/ramabmtr/go-barebone/app/errors"
	"github.com/ramabmtr/go-barebone/app/rest/handler"
	"github.com/ramabmtr/go-barebone/app/service/entity"
)

func SkipperByURLPath(paths ...string) func(c echo.Context) bool {
	return func(c echo.Context) bool {
		for i, _ := range paths {
			if strings.Contains(c.Request().URL.Path, paths[i]) {
				return true
			}
		}
		return false
	}
}

func Run(h *handler.Handler) {
	e := echo.New()

	e.Validator = config.GetValidator()

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, entity.MessageResponse{Message: "route not found"})
	}

	echo.MethodNotAllowedHandler = func(c echo.Context) error {
		return c.JSON(http.StatusMethodNotAllowed, entity.MessageResponse{Message: "method not allowed"})
	}

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if !c.Response().Committed {
			if c.Request().Method == http.MethodHead {
				err = c.NoContent(errors.ErrorToHTTPCode(err))
			} else {
				err = c.JSON(errors.ErrorToHTTPCode(err), entity.MessageResponse{Message: err.Error()})
			}
		}
	}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisableErrorHandler: true,
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper:   SkipperByURLPath("docs"),
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Println(fmt.Sprintf("[REQUEST] [%s] %v, status: %v", v.Method, v.URI, v.Status))
			return nil
		},
	}))

	RegisterRouter(e, h)

	go func() {
		err := e.Start(fmt.Sprintf(":%s", config.Conf.App.Port))
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), config.Conf.App.ShutdownTimeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
