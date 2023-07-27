package config

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
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
			log.Fatal("failed to connect to redis")
		}

		RedisConn = r
	})
}
