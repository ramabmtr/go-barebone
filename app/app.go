package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/ramabmtr/go-barebone/app/config"
	"github.com/ramabmtr/go-barebone/app/rest"
	restHandler "github.com/ramabmtr/go-barebone/app/rest/handler"
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
	rh := restHandler.InitHandler(uc)

	s := scheduler.New()
	s.Run()

	r := rest.New(rh)
	r.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), config.Conf.App.ShutdownTimeout)
	defer cancel()

	wg.Add(2)

	go func() {
		defer wg.Done()
		s.Stop(ctx)
	}()

	go func() {
		defer wg.Done()
		r.Stop(ctx)
	}()

	wg.Wait()
}
