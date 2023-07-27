package domain

import (
	"github.com/ramabmtr/go-barebone/app/service/domain/user"
)

type Domain struct {
	User      user.User
	UserCache user.Cache
}

func InitDomain() *Domain {
	return &Domain{
		User:      user.NewDomain(),
		UserCache: user.NewCacheDomain(),
	}
}
