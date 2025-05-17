package constant

const (
	LogTimeFormat = "2006-01-02T15:04:05.999Z07:00"
)

type LogFormat string

const (
	LogFormatJSON LogFormat = "json"
	LogFormatText LogFormat = "text"
)

type LogLevel string

const (
	LogLevelError LogLevel = "error"
	LogLevelInfo  LogLevel = "info"
	LogLevelDebug LogLevel = "debug"
)
