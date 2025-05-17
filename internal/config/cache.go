package config

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var (
	cacheOnce sync.Once
	cache     *redis.Client
)

func InitCacheConn() {
	cacheOnce.Do(func() {
		c := GetEnv().Cache
		if !c.IsActivated() {
			log.Warn().Msg("cache is not activated")
			return
		}

		cDB, err := strconv.Atoi(GetEnv().Cache.DB)
		if err != nil {
			cDB = 0
		}

		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", c.Host, c.Port),
			Password: GetEnv().Cache.Pass,
			DB:       cDB,
		})
		err = client.Ping(context.TODO()).Err()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to redis")
		}

		cache = client
	})
}

func GetCacheConn() *redis.Client {
	return cache
}
