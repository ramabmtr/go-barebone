package scheduler

import (
	"github.com/ramabmtr/go-barebone/internal/config"
	"github.com/ramabmtr/go-barebone/internal/scheduler/handler"
	"github.com/ramabmtr/go-barebone/internal/service"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

func RegisterRouter(c *cron.Cron, svc *service.Service) {
	dummyHandler := handler.NewDummyHandler(svc.Dummy)
	schedulerCfg := config.GetEnv().Scheduler

	_, err := c.AddFunc(schedulerCfg.DummyCronTab, CronFuncContextWrapper(dummyHandler.Dummy))
	if err != nil {
		log.Fatal().Err(err).Msg("error adding dummy scheduler")
	}
}
