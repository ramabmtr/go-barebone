package user

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ramabmtr/go-barebone/app/errors"
	"github.com/ramabmtr/go-barebone/app/service/entity"
	"github.com/ramabmtr/go-barebone/app/util/generator"
	"github.com/redis/go-redis/v9"
)

type cacheRedis struct {
	client *redis.Client
}

func NewUserCacheDomainRedis(client *redis.Client) Cache {
	return &cacheRedis{
		client: client,
	}
}

func (cd *cacheRedis) SetUser(ctx context.Context, u *entity.User, ttl time.Duration) error {
	if u.Username == "" {
		return fmt.Errorf("empty cache key using user.Username")
	}

	val, err := json.Marshal(u)
	if err != nil {
		return err
	}

	return cd.client.Set(ctx, generator.BuildCacheKey("user", u.Username), val, ttl).Err()
}

func (cd *cacheRedis) GetUser(ctx context.Context, username string) (*entity.User, error) {
	if username == "" {
		return nil, errors.ErrDataNotFound
	}

	val, err := cd.client.Get(ctx, generator.BuildCacheKey("user", username)).Bytes()
	if err == redis.Nil {
		return nil, errors.ErrDataNotFound
	}
	if err != nil {
		return nil, err
	}

	var u entity.User
	err = json.Unmarshal(val, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
