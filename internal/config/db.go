package config

import (
	"sync"
	"time"

	"github.com/ramabmtr/go-barebone/internal/constant"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbOnce sync.Once
	db     *gorm.DB
)

func getDBLogMode() logger.LogLevel {
	switch GetEnv().Log.GetLevel() {
	case constant.LogLevelDebug:
		return logger.Info
	case constant.LogLevelInfo:
		return logger.Warn
	default:
		return logger.Error
	}
}

func InitDBConn() {
	dbOnce.Do(func() {
		conf := gorm.Config{
			Logger:                 logger.Default.LogMode(getDBLogMode()),
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
			FullSaveAssociations:   false,
			NowFunc:                time.Now().UTC,
		}

		c := GetEnv().Database
		if !c.IsActivated() {
			log.Warn().Msg("database is not activated")
			return
		}
		var err error
		db, err = gorm.Open(postgres.Open(c.GetDSN()), &conf)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to database")
		}
	})
}

func GetDBConn() *gorm.DB {
	return db
}
