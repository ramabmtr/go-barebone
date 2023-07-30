package config

import (
	"os"
	"sync"
	"time"

	"github.com/ramabmtr/go-barebone/app/util"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

var (
	confOnce sync.Once

	Conf *config
)

const (
	DatabaseEngineInMemory = "inmemory"
	DatabaseEngineMySQL    = "mysql"

	CacheEngineInMemory = "inmemory"
	CacheEngineRedis    = "redis"
)

type config struct {
	FeatureFlag struct {
		EnableDocs bool `yaml:"enable_docs"`
	} `yaml:"feature_flag"`

	App struct {
		Port            string        `yaml:"port"`
		ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
		JWT             struct {
			Secret      string        `yaml:"secret"`
			ExpiredTime time.Duration `yaml:"expired_time"`
		} `yaml:"jwt"`
	} `yaml:"app"`

	Database struct {
		Engine string `yaml:"engine"`
		Config struct {
			MySQL struct {
				Host   string `yaml:"host"`
				Port   string `yaml:"port"`
				User   string `yaml:"user"`
				Pass   string `yaml:"pass"`
				DBName string `yaml:"db_name"`
			} `yaml:"mysql"`
		} `yaml:"config"`
	} `yaml:"database"`

	Cache struct {
		Engine string `yaml:"engine"`
		Config struct {
			Redis struct {
				Host string `yaml:"host"`
				Port string `yaml:"port"`
				Pass string `yaml:"pass"`
				DB   int    `yaml:"db"`
			} `yaml:"redis"`
		} `yaml:"config"`
	} `yaml:"cache"`

	Scheduler []struct {
		Name    string `yaml:"name"`
		Crontab string `yaml:"crontab"`
	} `yaml:"scheduler"`
}

func InitConf(confPath string) {
	confOnce.Do(func() {
		c := &config{}

		yamlFile, err := os.ReadFile(confPath)
		if err != nil {
			log.Fatal().Err(err).Msg("read yaml failed")
		}

		err = yaml.Unmarshal(yamlFile, c)
		if err != nil {
			log.Fatal().Err(err).Msg("unmarshal yaml failed")
		}

		// check enum value
		if c.Database.Engine == "" || !util.StrIn(c.Database.Engine, DatabaseEngineInMemory, DatabaseEngineMySQL) {
			log.Fatal().Msg("database.engine is not set or not supported")
		}

		if c.Cache.Engine == "" || !util.StrIn(c.Cache.Engine, CacheEngineInMemory, CacheEngineRedis) {
			log.Fatal().Msg("cache.engine is not set or not supported")
		}

		Conf = c
	})
}
