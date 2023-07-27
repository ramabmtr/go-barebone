package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ramabmtr/go-barebone/app/config"
	"github.com/ramabmtr/go-barebone/app/errors"
	userDom "github.com/ramabmtr/go-barebone/app/service/domain/user"
	"github.com/ramabmtr/go-barebone/app/service/entity"
	"github.com/ramabmtr/go-barebone/app/util/generator"
)

type authUseCase struct {
	userDomain userDom.User
}

func NewAuthUseCase(userDomain userDom.User) Auth {
	return &authUseCase{
		userDomain: userDomain,
	}
}

func (uc *authUseCase) jwtGen(user *entity.User) (string, error) {
	now := time.Now()
	claims := &entity.JWTCustomClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(config.Conf.App.JWT.ExpiredTime)),
			Subject:   generator.RandomString(8, generator.Alphanumeric),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Conf.App.JWT.Secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (uc *authUseCase) Register(ctx context.Context, p *entity.UserRegisterParam) (*entity.UserLoginResponse, error) {
	currentUser := entity.User{
		Username: p.Username,
	}

	err := uc.userDomain.Get(ctx, &currentUser)
	if err != nil && err != errors.ErrDataNotFound {
		return nil, err
	}

	if !currentUser.CreatedAt.IsZero() {
		return nil, errors.ErrUserAlreadyRegistered
	}

	now := time.Now().UTC()

	user := entity.User{
		Username:  p.Username,
		CreatedAt: now,
	}

	err = uc.userDomain.Create(ctx, &user)
	if err != nil {
		return nil, err
	}

	token, err := uc.jwtGen(&user)
	if err != nil {
		return nil, err
	}

	return &entity.UserLoginResponse{Token: token}, nil
}

func (uc *authUseCase) Login(ctx context.Context, p *entity.UserLoginParam) (*entity.UserLoginResponse, error) {
	user := entity.User{
		Username: p.Username,
	}

	err := uc.userDomain.Get(ctx, &user)
	if err != nil {
		return nil, err
	}

	token, err := uc.jwtGen(&user)
	if err != nil {
		return nil, err
	}

	return &entity.UserLoginResponse{Token: token}, nil
}
