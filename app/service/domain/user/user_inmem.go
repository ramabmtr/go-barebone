package user

import (
	"context"
	"sync"

	"github.com/ramabmtr/go-barebone/app/errors"
	"github.com/ramabmtr/go-barebone/app/service/entity"
)

type userInMem struct {
	m    sync.Mutex
	user map[string]*entity.User
}

func NewUserDomainInMem() User {
	return &userInMem{
		user: make(map[string]*entity.User),
	}
}

func (ud *userInMem) Create(_ context.Context, u *entity.User) error {
	ud.m.Lock()
	defer ud.m.Unlock()

	if ud.user[u.Username] != nil {
		return errors.ErrUserAlreadyRegistered
	}

	ud.user[u.Username] = u

	return nil
}

func (ud *userInMem) Get(_ context.Context, u *entity.User) error {
	if u.Username != "" && ud.user[u.Username] != nil {
		*u = *ud.user[u.Username]
		return nil
	}

	return errors.ErrDataNotFound
}
