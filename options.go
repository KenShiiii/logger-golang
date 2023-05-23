package splog

import "os"

type loggerOptions struct {
	logLevel int
	logFile  *os.File
}

func NewLoggerOptions() *loggerOptions {
	opts := &loggerOptions{
		logLevel: 0,
		logFile:  nil,
	}
	return opts
}

func (l *loggerOptions) SetLogLevel(lv int) {
	switch {
	case lv < 0:
		l.logLevel = 0
	case lv > 3:
		l.logLevel = 3
	default:
		l.logLevel = lv
	}
}
