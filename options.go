package gologger

import "os"

type loggerOptions struct {
	logLevel int
	logFile  os.File
}
