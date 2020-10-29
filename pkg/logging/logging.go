package logging

import "github.com/Matt-Gleich/logoru"

type LogLevel string

const (
	LogLevelCritical LogLevel = "critical"
	LogLevelInfo     LogLevel = "info"
)

func Log(msg interface{}, level LogLevel) {
	switch level {
	case LogLevelCritical:
		logoru.Critical(msg)
	case LogLevelInfo:
		logoru.Info(msg)
	}
}

func Info(msg interface{}) {
	Log(msg, LogLevelInfo)
}

func Critical(msg interface{}) {
	Log(msg, LogLevelCritical)
}
