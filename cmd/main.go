package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"

	"github.com/ramabmtr/go-barebone/internal/api"
	"github.com/ramabmtr/go-barebone/internal/config"
	"github.com/ramabmtr/go-barebone/internal/repository"
	"github.com/ramabmtr/go-barebone/internal/scheduler"
	"github.com/ramabmtr/go-barebone/internal/service"
)

func main() {
	var envPath string

	flag.StringVar(&envPath, "env", ".env", "path to env file")
	flag.Parse()

	config.InitEnv(envPath)
	config.InitLog()
	config.InitDBConn()
	config.InitCacheConn()

	repo := repository.InitRepository()
	svc := service.InitService(repo)

	apiRunner := api.New(svc)
	apiRunner.Run()

	schRunner := scheduler.New(svc)
	schRunner.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), config.GetEnv().Server.ShutdownTimeout)
	defer cancel()

	wg.Add(2)

	go func() {
		defer wg.Done()
		apiRunner.Stop(ctx)
	}()

	go func() {
		defer wg.Done()
		schRunner.Stop(ctx)
	}()

	wg.Wait()
}
