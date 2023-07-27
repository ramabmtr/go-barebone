package auth

import (
	"context"

	"github.com/ramabmtr/go-barebone/app/service/entity"
)

type Auth interface {
	Register(ctx context.Context, p *entity.UserRegisterParam) (*entity.UserLoginResponse, error)
	Login(ctx context.Context, p *entity.UserLoginParam) (*entity.UserLoginResponse, error)
}
