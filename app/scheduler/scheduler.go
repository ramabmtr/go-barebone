package scheduler

import (
	"context"

	"github.com/labstack/gommon/random"
	"github.com/ramabmtr/go-barebone/app/config"
	"github.com/ramabmtr/go-barebone/app/util/appctx"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
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
			log.Error().Err(err).Msgf("error when processing scheduler with ID: %s", rid)
		}
	}
}

func (s *Scheduler) Run() {
	for _, sc := range config.Conf.Scheduler {
		fn, ok := mapCronFunc[sc.Name]
		if !ok {
			log.Warn().Str("name", sc.Name).Msg("no route for scheduler")
			continue
		}

		_, err := s.c.AddFunc(sc.Crontab, fn)
		if err != nil {
			log.Fatal().Err(err).Str("name", sc.Name).Msg("error adding scheduler")
		}

		log.Info().Str("name", sc.Name).Msg("scheduler registered")
	}

	s.c.Start()
}

func (s *Scheduler) Stop(ctx context.Context) {
	defer log.Info().Msg("scheduler stopped")
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
