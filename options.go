package splog

import "os"

type loggerOptions struct {
	logLevel int
	logFile  os.File
}
