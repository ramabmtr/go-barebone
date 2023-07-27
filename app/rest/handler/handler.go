package handler

import "github.com/ramabmtr/go-barebone/app/service/usecase"

type Handler struct {
	Ping *ping
	User *user
}

func InitHandler(uc *usecase.UseCase) *Handler {
	return &Handler{
		Ping: newPingHandler(),
		User: newUserHandler(uc.Auth),
	}
}
