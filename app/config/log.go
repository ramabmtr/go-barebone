package config

import (
	"os"

	"github.com/ramabmtr/go-barebone/app/util/appctx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	logTimeFormat = "2006-01-02T15:04:05.999Z07:00"

	LogFormatJSON = "json"
	LogFormatText = "text"

	LogLevelError = "error"
	LogLevelInfo  = "info"
	LogLevelDebug = "debug"
)

func InitLog() {
	zerolog.TimeFieldFormat = logTimeFormat

	setLogFormat()
	setLogLevel()
}

func setLogFormat() {
	var l zerolog.Logger
	switch Conf.App.Log.Format {
	case LogFormatJSON:
		l = log.Logger
	case LogFormatText:
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: logTimeFormat}
		l = zerolog.New(output)
	default:
		l = log.Logger
	}

	log.Logger = l.With().Caller().Timestamp().Logger().Hook(tracingHook{})
}

func setLogLevel() {
	switch Conf.App.Log.Level {
	case LogLevelError:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case LogLevelInfo:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case LogLevelDebug:
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

	authInfo := appctx.GetAuthInfo(ctx)
	if authInfo != nil {
		e.Str("username", authInfo.Username)
	}
}
