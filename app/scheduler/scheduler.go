package scheduler

import (
	"context"
	"log"

	"github.com/labstack/gommon/random"
	"github.com/ramabmtr/go-barebone/app/config"
	"github.com/ramabmtr/go-barebone/app/util/appctx"
	"github.com/robfig/cron/v3"
)

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

func Run() {
	c := cron.New(cron.WithParser(cron.NewParser(
		cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow,
	)))

	for _, s := range config.Conf.Scheduler {
		fn, ok := mapCronFunc[s.Name]
		if !ok {
			log.Printf("no route for scheduler \"%s\"\n", s.Name)
			continue
		}

		_, err := c.AddFunc(s.Crontab, fn)
		if err != nil {
			log.Fatalf("error adding scheduler \"%s\". %s\n", s.Name, err.Error())
		}

		log.Printf("scheduler \"%s\" registered\n", s.Name)
	}

	c.Start()
}
