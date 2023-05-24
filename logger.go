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
	Log(level int, args ...any)

	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
}

type logger struct {
	logLevel uint // LogLevel specifies the minimum log level required for a log message to be printed.

	filePath string   // FilePath specifies the path to the log file.
	file     *os.File // file is the file handle for the log file.
	stdout   bool
}

const (
	LevelDebug uint = 0
	LevelInfo  uint = 1
	LevelWarn  uint = 2
	LevelErr   uint = 3
)

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
		stdout:   o.stdout,
	}

	return l
}

// Log prints a log message with the given log level and message arguments.
// If the log level is less than the logger's LogLevel, the message is not printed.
func (l *logger) Log(level int, args ...interface{}) {
	if uint(level) < l.logLevel {
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

	_, file, line, ok := runtime.Caller(1)
	if !ok {

		file = ""
		line = 0
	}

	fileName := path.Base(file)
	// shottFuncName := path.Base(runtime.FuncForPC(funcName).Name())
	logTime := time.Now().Local().Format("2006-01-02 15:04:05.000 UTC-0700")

	fmt.Printf("[%s] %s[%s:%d] %s \n", logTime, prefix, fileName, line, fmt.Sprint(args...))

	// Replace escape characters with an empty string
	re := regexp.MustCompile(`\033\[[0-9;]*[a-zA-Z]`)
	prefix = re.ReplaceAllString(prefix, "")
	argsStr := re.ReplaceAllString(fmt.Sprint(args...), "")

	// write to log file
	if l.file != nil {
		fmt.Fprintf(l.file, "[%s] %s[%s:%d] %s \n", logTime, prefix, fileName, line, argsStr)
	}

}

// writeLog is a helper function used to write log messages to both stdout and a log file.
// It takes a log message prefix and variadic arguments for the log message content.
// It retrieves information about the caller's file name and line number using the runtime package.
// The log message is formatted with the log time, prefix, file name, line number, and log message content.
// If enabled, it prints the log message to stdout.
// It also writes the log message to the log file, after removing ANSI escape characters.
func (l *logger) writeLog(prefix string, args ...any) {
	// Retrieve file name and line number of the caller
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = ""
		line = 0
	}

	fileName := path.Base(file) // Extract file name from the file path
	logTime := time.Now().Local().Format("2006-01-02 15:04:05.000 UTC-0700")

	if l.stdout {
		// Print log message to stdout
		fmt.Printf("[%s] %s[%s:%d] %s \n", logTime, prefix, fileName, line, fmt.Sprint(args...))
	}

	// Replace escape characters with an empty string
	re := regexp.MustCompile(`\033\[[0-9;]*[a-zA-Z]`)
	prefix = re.ReplaceAllString(prefix, "")
	argsStr := re.ReplaceAllString(fmt.Sprint(args...), "")

	if l.file != nil {
		// Write log message to the log file
		fmt.Fprintf(l.file, "[%s] %s[%s:%d] %s \n", logTime, prefix, fileName, line, argsStr)
	}
}

// Debug logs messages with the "DEBU" prefix if the log level is set to debug.
// The messages are written to the log file and, if enabled, printed to stdout.
func (l *logger) Debug(args ...any) {
	if l.logLevel > LevelDebug {
		return
	}
	prefix := "\033[38;5;16;48;5;253mDEBU\033[0m "

	l.writeLog(prefix, args...)
}

// Info logs messages with the "INFO" prefix if the log level is set to info or below.
// The messages are written to the log file and, if enabled, printed to stdout.
func (l *logger) Info(args ...any) {
	if l.logLevel > LevelInfo {
		return
	}
	prefix := "\033[48;5;57mINFO\033[0m "

	l.writeLog(prefix, args...)
}

// Warn logs messages with the "WARN" prefix if the log level is set to warn or below.
// The messages are written to the log file and, if enabled, printed to stdout.
func (l *logger) Warn(args ...any) {
	if l.logLevel > LevelWarn {
		return
	}
	prefix := "\033[38;5;16;48;5;226mWARN\033[0m "

	l.writeLog(prefix, args...)
}

// Error logs messages with the "ERRO" prefix.
// The messages are written to the log file and, if enabled, printed to stdout.
func (l *logger) Error(args ...any) {
	prefix := "\033[38;5;15;48;5;160mERRO\033[0m "

	l.writeLog(prefix, args...)
}

// Fatal logs messages with the "FATL" prefix and exits the program with a status of 1.
// The messages are written to the log file and, if enabled, printed to stdout.
func (l *logger) Fatal(args ...any) {
	prefix := "\033[38;5;15;48;5;196;5mFATL\033[0m "

	l.writeLog(prefix, args...)
	os.Exit(1)
}
