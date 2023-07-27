package usecase

import (
	"github.com/ramabmtr/go-barebone/app/service/domain"
	"github.com/ramabmtr/go-barebone/app/service/usecase/auth"
)

type UseCase struct {
	Auth auth.Auth
}

func InitUseCase(d *domain.Domain) *UseCase {
	return &UseCase{
		Auth: auth.NewAuthUseCase(d.User),
	}
}
