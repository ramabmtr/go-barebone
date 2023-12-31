package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	mysqlOnce sync.Once

	MySQLConn *gorm.DB
)

func InitMySQLConn() {
	mysqlOnce.Do(func() {
		conf := gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Info),
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
			FullSaveAssociations:   false,
			NowFunc:                time.Now().UTC,
		}

		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
			Conf.Database.Config.MySQL.User,
			Conf.Database.Config.MySQL.Pass,
			Conf.Database.Config.MySQL.Host,
			Conf.Database.Config.MySQL.Port,
			Conf.Database.Config.MySQL.DBName,
		)
		db, err := gorm.Open(mysql.Open(dsn), &conf)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to mysql")
		}

		MySQLConn = db
	})
}
