package scheduler

import (
	"context"
	"log"

	"github.com/labstack/gommon/random"
	"github.com/ramabmtr/go-barebone/app/config"
	"github.com/ramabmtr/go-barebone/app/util/appctx"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	c *cron.Cron
}

func New() *Scheduler {
	return &Scheduler{
		c: cron.New(cron.WithParser(cron.NewParser(
			cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow,
		))),
	}
}

func CronFuncContextWrapper(f func(context.Context) error) func() {
	return func() {
		ctx := context.TODO()
		rid := random.String(32)
		ctx = appctx.SetRequestID(ctx, rid)

		err := f(ctx)
		if err != nil {
			log.Printf("error when processing scheduler with ID: %s. Err: %s", rid, err.Error())
		}
	}
}

func (s *Scheduler) Run() {
	for _, sc := range config.Conf.Scheduler {
		fn, ok := mapCronFunc[sc.Name]
		if !ok {
			log.Printf("no route for scheduler \"%s\"\n", sc.Name)
			continue
		}

		_, err := s.c.AddFunc(sc.Crontab, fn)
		if err != nil {
			log.Fatalf("error adding scheduler \"%s\". %s\n", sc.Name, err.Error())
		}

		log.Printf("scheduler \"%s\" registered\n", sc.Name)
	}

	s.c.Start()
}

func (s *Scheduler) Stop(ctx context.Context) {
	defer log.Println("scheduler stopped")
	ctxS := s.c.Stop()

	for {
		select {
		case <-ctxS.Done():
			return
		case <-ctx.Done():
			return
		}
	}
}
