package logging

import "github.com/Matt-Gleich/logoru"

type LogLevel string

const (
	LogLevelCritical LogLevel = "critical"
)

func Log(msg interface{}, level LogLevel) {
	switch level {
	case LogLevelCritical:
		logoru.Critical(msg)
	}
}
