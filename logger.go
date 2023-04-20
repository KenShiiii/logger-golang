package logger

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

// Logger represents a logger instance with a specific log level.
type Logger struct {
	// LogLevel specifies the minimum log level required for a log message to be printed.
	LogLevel int

	// FilePath specifies the path to the log file.
	FilePath string

	// file is the file handle for the log file.
	File *os.File
}

// Log prints a log message with the given log level and message arguments.
// If the log level is less than the logger's LogLevel, the message is not printed.
func (l *Logger) Log(level int, args ...interface{}) {
	if level >= l.LogLevel {
		var prefix string
		switch level {
		case 0:
			// Debug level
			prefix = "[DEBUG] "
		case 1:
			// Info level
			prefix = "[\033[34mINFO\033[0m] "
		case 2:
			// Warning level
			prefix = "[\033[93mWARNING\033[0m] "
		case 3:
			// Error level
			prefix = "[\033[91mERROR\033[0m] "
		default:
			// Unknown level
			prefix = "[UNKNOWN LEVEL] "
		}
		log.Printf("%s%s", prefix, fmt.Sprint(args...))

		// Replace escape characters with an empty string
		re := regexp.MustCompile(`\033\[[0-9;]*[a-zA-Z]`)
		prefix = re.ReplaceAllString(prefix, "")
		argsStr := re.ReplaceAllString(fmt.Sprint(args...), "")

		logTime := time.Now().Local().Format("2006-01-02 15:04:05 UTC-0700")

		// write to log file
		if l.File != nil {
			fmt.Fprintf(l.File, "\n%s %s%s", logTime, prefix, argsStr)
		}
	}
}

// Close closes the log file.
func (l *Logger) Close() {
	if l.File != nil {
		l.File.Close()
		l.File = nil
	}
}
