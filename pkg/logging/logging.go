package logging

import "github.com/Matt-Gleich/logoru"

type LogLevel string

const (
	LogLevelCritical LogLevel = "critical"
	LogLevelInfo     LogLevel = "info"
	LogLevelError    LogLevel = "error"
)

func Log(msg interface{}, level LogLevel) {
	switch level {
	case LogLevelCritical:
		logoru.Critical(msg)
	case LogLevelInfo:
		logoru.Info(msg)
	case LogLevelError:
		logoru.Error(msg)
	}
}

func Info(msg interface{}) {
	Log(msg, LogLevelInfo)
}

func Critical(msg interface{}) {
	Log(msg, LogLevelCritical)
}

func Error(msg interface{}) {
	Log(msg, LogLevelError)
}
