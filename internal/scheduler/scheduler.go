package scheduler

import (
	"context"

	"github.com/ramabmtr/go-barebone/internal/lib/appctx"
	"github.com/ramabmtr/go-barebone/internal/lib/generator"
	"github.com/ramabmtr/go-barebone/internal/service"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

type Scheduler struct {
	c   *cron.Cron
	svc *service.Service
}

func New(svc *service.Service) *Scheduler {
	return &Scheduler{
		c: cron.New(cron.WithParser(cron.NewParser(
			cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow,
		))),
		svc: svc,
	}
}

func CronFuncContextWrapper(f func(context.Context) error) func() {
	return func() {
		ctx := context.TODO()
		rid := generator.RandomString(32, generator.Alphanumeric)
		ctx = appctx.SetRequestID(ctx, rid)

		err := f(ctx)
		if err != nil {
			log.Error().Err(err).Msgf("error when processing scheduler with ID: %s", rid)
		}
	}
}

func (s *Scheduler) Run() {
	RegisterRouter(s.c, s.svc)
	s.c.Start()
	log.Info().Msg("scheduler started")
}

func (s *Scheduler) Stop(ctx context.Context) {
	ctxS := s.c.Stop()

	for {
		select {
		case <-ctxS.Done():
			log.Info().Msg("scheduler stopped gracefully")
			return
		case <-ctx.Done():
			log.Info().Msg("scheduler stopped because of timeout")
			return
		}
	}
}
