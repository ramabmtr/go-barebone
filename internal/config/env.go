package config

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/ramabmtr/go-barebone/internal/constant"
	"github.com/rs/zerolog/log"
)

var (
	envOnce sync.Once
	env     *Env
)

type Env struct {
	Server    *ServerEnv    `envconfig:"SERVER"`
	Log       *LogEnv       `envconfig:"LOG"`
	Database  *DatabaseEnv  `envconfig:"DATABASE"`
	Cache     *CacheEnv     `envconfig:"CACHE"`
	Scheduler *SchedulerEnv `envconfig:"SCHEDULER"`
}

type ServerEnv struct {
	Port            string        `envconfig:"PORT" default:"8080"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"5s"`
}

type LogEnv struct {
	Format string `envconfig:"FORMAT" default:"json"`
	Level  string `envconfig:"LEVEL" default:"error"`
}

func (l *LogEnv) GetLevel() constant.LogLevel {
	switch strings.ToLower(GetEnv().Log.Level) {
	case "error":
		return constant.LogLevelError
	case "debug":
		return constant.LogLevelDebug
	default:
		return constant.LogLevelInfo
	}
}

func (l *LogEnv) GetFormat() constant.LogFormat {
	switch strings.ToLower(GetEnv().Log.Format) {
	case "text":
		return constant.LogFormatText
	default:
		return constant.LogFormatJSON
	}
}

type DatabaseEnv struct {
	Host    string `envconfig:"HOST"`
	Port    string `envconfig:"PORT"`
	User    string `envconfig:"USER"`
	Pass    string `envconfig:"PASS"`
	DBName  string `envconfig:"NAME"`
	SSLMode string `envconfig:"SSL_MODE"`
}

// IsActivated is used to check if the database is used or not based on all the env variables
// if all the env variables are empty except Pass, then the database is not used
func (d *DatabaseEnv) IsActivated() bool {
	if d.Host == "" && d.Port == "" && d.User == "" && d.DBName == "" && d.SSLMode == "" {
		return false
	}
	return true
}

// GetDSN is used to build DSN based on env variables
func (d *DatabaseEnv) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Pass, d.DBName, d.SSLMode,
	)
}

type CacheEnv struct {
	Host string `envconfig:"HOST"`
	Port string `envconfig:"PORT"`
	Pass string `envconfig:"PASS"`
	DB   string `envconfig:"DB"`
}

// IsActivated is used to check if the database is used or not based on all the env variables
// if all the env variables are empty except Pass and DB, then the database is not used
func (c *CacheEnv) IsActivated() bool {
	if c.Host == "" && c.Port == "" {
		return false
	}
	return true
}

type SchedulerEnv struct {
	DummyCronTab string `envconfig:"DUMMY_CRON_TAB" default:"* * * * *"`
}

func InitEnv(filenames ...string) {
	envOnce.Do(func() {
		err := godotenv.Load(filenames...)
		if err != nil {
			log.Warn().Err(err).Msg(".env file not found, using environment variables")
		}

		env = &Env{}

		err = envconfig.Process("", env)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to process env")
		}
	})
}

func GetEnv() *Env {
	return env
}
