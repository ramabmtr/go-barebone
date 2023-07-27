package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ramabmtr/go-barebone/app/errors"
	"github.com/ramabmtr/go-barebone/app/service/entity"
	"github.com/ramabmtr/go-barebone/app/service/usecase/auth"
)

type user struct {
	authUC auth.Auth
}

func newUserHandler(authUC auth.Auth) *user {
	return &user{
		authUC: authUC,
	}
}

// Register :
// @description Register the user to the system
// @tags Auth
// @accept json
// @produce json
// @param data body entity.UserRegisterParam true "payload"
// @success 200 {object} entity.UserLoginResponse
// @failure 400 {object} entity.MessageResponse
// @failure 409 {object} entity.MessageResponse
// @router /register [post]
func (h *user) Register(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(entity.UserRegisterParam)

	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.MessageResponse{Message: err.Error()})
	}

	err = c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.MessageResponse{Message: err.Error()})
	}

	loginInfo, err := h.authUC.Register(ctx, req)
	if err != nil {
		return c.JSON(errors.ErrorToHTTPCode(err), entity.MessageResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, loginInfo)
}

// Login :
// @description Login the user to the system
// @tags Auth
// @accept json
// @produce json
// @param data body entity.UserLoginParam true "payload"
// @success 200 {object} entity.UserLoginResponse
// @failure 400 {object} entity.MessageResponse
// @failure 409 {object} entity.MessageResponse
// @router /login [post]
func (h *user) Login(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(entity.UserLoginParam)

	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.MessageResponse{Message: err.Error()})
	}

	err = c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.MessageResponse{Message: err.Error()})
	}

	loginInfo, err := h.authUC.Login(ctx, req)
	if err != nil {
		return c.JSON(errors.ErrorToHTTPCode(err), entity.MessageResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, loginInfo)
}
