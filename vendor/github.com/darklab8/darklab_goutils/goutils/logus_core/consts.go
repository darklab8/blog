package logus_core

type LogLevel string

func (l LogLevel) ToStr() string { return string(l) }

const (
	LEVEL_DEBUG LogLevel = "DEBUG"
	LEVEL_INFO  LogLevel = "INFO"
	LEVEL_WARN  LogLevel = "WARN"
	LEVEL_ERROR LogLevel = "ERROR"
)
