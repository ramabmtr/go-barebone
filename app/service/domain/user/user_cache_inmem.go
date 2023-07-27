package user

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ramabmtr/go-barebone/app/errors"
	"github.com/ramabmtr/go-barebone/app/service/entity"
)

type cacheInMem struct {
	m    sync.Mutex
	user map[string]*userTTL
}

type userTTL struct {
	user *entity.User
	ttl  time.Time
}

func NewUserCacheDomainInMem() Cache {
	return &cacheInMem{
		user: make(map[string]*userTTL),
	}
}

func (cd *cacheInMem) SetUser(_ context.Context, u *entity.User, ttl time.Duration) error {
	if u.Username == "" {
		return fmt.Errorf("empty cache key using user.Username")
	}

	ut := userTTL{
		user: u,
		ttl:  time.Now().UTC().Add(ttl),
	}

	cd.m.Lock()
	defer cd.m.Unlock()

	cd.user[u.Username] = &ut

	return nil
}

func (cd *cacheInMem) GetUser(_ context.Context, username string) (*entity.User, error) {
	if username == "" {
		return nil, errors.ErrDataNotFound
	}

	u, ok := cd.user[username]
	if !ok || u == nil {
		return nil, errors.ErrDataNotFound
	}

	if time.Now().UTC().After(u.ttl) {
		return nil, errors.ErrDataNotFound
	}

	return u.user, nil
}
