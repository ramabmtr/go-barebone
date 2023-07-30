package config

import (
	"os"

	"github.com/ramabmtr/go-barebone/app/util/appctx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const logTimeFormat = "2006-01-02T15:04:05.999Z07:00"

func InitLog() {
	zerolog.TimeFieldFormat = logTimeFormat
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: logTimeFormat}
	log.Logger = zerolog.New(output).
		With().Caller().Timestamp().
		Logger().
		Hook(tracingHook{})
}

type tracingHook struct{}

func (h tracingHook) Run(e *zerolog.Event, _ zerolog.Level, _ string) {
	ctx := e.GetCtx()

	rid := appctx.GetRequestID(ctx)
	if rid != "" {
		e.Str("request_id", rid)
	}

	authInfo := appctx.GetAuthInfo(ctx)
	if authInfo != nil {
		e.Str("username", authInfo.Username)
	}
}
