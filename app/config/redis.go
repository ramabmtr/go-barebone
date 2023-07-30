package config

import (
	"context"
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var (
	redisOnce sync.Once

	RedisConn *redis.Client
)

func InitRedisConn() {
	redisOnce.Do(func() {
		r := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", Conf.Cache.Config.Redis.Host, Conf.Cache.Config.Redis.Port),
			Password: Conf.Cache.Config.Redis.Pass,
			DB:       Conf.Cache.Config.Redis.DB,
		})
		err := r.Ping(context.TODO()).Err()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to redis")
		}

		RedisConn = r
	})
}
