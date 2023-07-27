package user

import (
	"context"
	"time"

	"github.com/ramabmtr/go-barebone/app/config"
	"github.com/ramabmtr/go-barebone/app/service/entity"
)

type User interface {
	Create(ctx context.Context, u *entity.User) error
	Get(ctx context.Context, u *entity.User) error
}

type Cache interface {
	SetUser(ctx context.Context, u *entity.User, ttl time.Duration) error
	GetUser(ctx context.Context, username string) (*entity.User, error)
}

func NewDomain() User {
	switch config.Conf.Database.Engine {
	case config.DatabaseEngineMySQL:
		return NewUserDomainMySQL(config.MySQLConn)
	default:
		return NewUserDomainInMem()
	}
}

func NewCacheDomain() Cache {
	switch config.Conf.Cache.Engine {
	case config.CacheEngineRedis:
		return NewUserCacheDomainRedis(config.RedisConn)
	default:
		return NewUserCacheDomainInMem()
	}
}
