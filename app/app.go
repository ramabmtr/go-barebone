package app

import (
	"log"

	"github.com/ramabmtr/go-barebone/app/config"
	"github.com/ramabmtr/go-barebone/app/rest"
	"github.com/ramabmtr/go-barebone/app/rest/handler"
	"github.com/ramabmtr/go-barebone/app/scheduler"
	"github.com/ramabmtr/go-barebone/app/service/domain"
	"github.com/ramabmtr/go-barebone/app/service/entity"
	"github.com/ramabmtr/go-barebone/app/service/usecase"
)

func Run() {
	// if db engine is mysql, prepare the connection and schema
	if config.Conf.Database.Engine == config.DatabaseEngineMySQL {
		config.InitMySQLConn()

		err := config.MySQLConn.AutoMigrate(
			&entity.User{},
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	if config.Conf.Cache.Engine == config.CacheEngineRedis {
		config.InitRedisConn()
	}

	dom := domain.InitDomain()
	uc := usecase.InitUseCase(dom)
	restHandler := handler.InitHandler(uc)

	scheduler.Run()
	rest.Run(restHandler)
}
