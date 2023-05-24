package splog

import "os"

type loggerOptions struct {
	logLevel uint
	logFile  *os.File
	stdout   bool
}

func NewLoggerOptions() *loggerOptions {
	opts := &loggerOptions{
		logLevel: 0,
		logFile:  nil,
		stdout:   true,
	}
	return opts
}

func (l *loggerOptions) SetLogLevel(lv int) {
	switch {
	case lv < 0:
		l.logLevel = LevelDebug
	case lv > 4:
		l.logLevel = LevelFatal
	default:
		l.logLevel = uint(lv)
	}
}

// EnableStdoutLogging enables or disables logging to standard output (stdout).
// If enable is set to true, log messages will be written to stdout in addition
// to the log file. If enable is set to false, log messages will only be written
// to the log file.
func (l *loggerOptions) EnableStdoutLogging(enable bool) {
	l.stdout = enable
}
