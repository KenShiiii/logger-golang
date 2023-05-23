package splog

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"time"
)

// Logger is the interface difinition for a simple logger as used by this library.
type Logger interface {
	Close()
	Log(level int, args ...interface{})
}

type logger struct {
	logLevel int // LogLevel specifies the minimum log level required for a log message to be printed.

	filePath string // FilePath specifies the path to the log file.

	file *os.File // file is the file handle for the log file.
}

// Close closes the log file.
func (l *logger) Close() {
	if l.file != nil {
		l.file.Close()
		l.file = nil
	}
}

func NewLogger(o *loggerOptions) Logger {
	l := &logger{
		logLevel: o.logLevel,
		file:     o.logFile,
	}

	return l
}

// Log prints a log message with the given log level and message arguments.
// If the log level is less than the logger's LogLevel, the message is not printed.
func (l *logger) Log(level int, args ...interface{}) {
	if level < l.logLevel {
		return
	}
	var prefix string
	switch level {
	case 0:
		// Debug level
		prefix = "\033[38;5;16;48;5;253mDEBU\033[0m "
	case 1:
		// Info level
		prefix = "\033[48;5;57mINFO\033[0m "
	case 2:
		// Warning level
		prefix = "\033[38;5;16;48;5;226mWARN\033[0m "
	case 3:
		// Error level
		prefix = "\033[38;5;15;48;5;160mERRO\033[0m "
	default:
		// Unknown level
		prefix = "\033[38;5;15;48;5;82mUNKNOWN LEVEL\033[0m "
	}

	// empty line to seperate log
	// fmt.Println()
	_, file, line, ok := runtime.Caller(1)
	if !ok {

		file = ""
		line = 0
	}

	fileName := path.Base(file)
	// shottFuncName := path.Base(runtime.FuncForPC(funcName).Name())
	logTime := time.Now().Local().Format("2006-01-02 15:04:05.000 UTC-0700")

	fmt.Printf("[%s] %s[%s:%d] %s \n", logTime, prefix, fileName, line, fmt.Sprint(args...))

	// log.Printf("%s%s", prefix, fmt.Sprint(args...))

	// Replace escape characters with an empty string
	re := regexp.MustCompile(`\033\[[0-9;]*[a-zA-Z]`)
	prefix = re.ReplaceAllString(prefix, "")
	argsStr := re.ReplaceAllString(fmt.Sprint(args...), "")

	// write to log file
	if l.file != nil {
		fmt.Fprintf(l.file, "[%s] %s[%s:%d] %s \n", logTime, prefix, fileName, line, argsStr)
	}

}
