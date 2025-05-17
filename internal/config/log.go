package config

import (
	"os"

	"github.com/ramabmtr/go-barebone/internal/constant"
	"github.com/ramabmtr/go-barebone/internal/lib/appctx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLog() {
	zerolog.TimeFieldFormat = constant.LogTimeFormat

	setLogFormat()
	setLogLevel()
}

func setLogFormat() {
	var l zerolog.Logger
	switch GetEnv().Log.GetFormat() {
	case constant.LogFormatJSON:
		l = log.Logger
	case constant.LogFormatText:
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: constant.LogTimeFormat}
		l = zerolog.New(output)
	default:
		l = log.Logger
	}

	log.Logger = l.With().Caller().Timestamp().Logger().Hook(tracingHook{})
}

func setLogLevel() {
	switch GetEnv().Log.GetLevel() {
	case constant.LogLevelError:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case constant.LogLevelInfo:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case constant.LogLevelDebug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

type tracingHook struct{}

func (h tracingHook) Run(e *zerolog.Event, _ zerolog.Level, _ string) {
	ctx := e.GetCtx()

	rid := appctx.GetRequestID(ctx)
	if rid != "" {
		e.Str("request_id", rid)
	}

	userInfo := appctx.GetUserInfo(ctx)
	if userInfo != nil {
		e.Str("username", userInfo.Username)
	}
}
