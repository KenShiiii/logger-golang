package splog

import (
	"log"
	"os"
)

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

// SetLogLevel sets the log level for the logger.
// The log level determines which log messages are recorded based on their severity.
// The input lv should be an integer value representing the desired log level:
//   - Values less than 0 will set the log level to Debug.
//   - Values greater than 4 will set the log level to Fatal.
//   - Values between 0 and 4 (inclusive) will set the log level accordingly.
//   - 0: Debug
//   - 1: Info
//   - 2: Warning
//   - 3: Error
//   - 4: Fatal
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

// SetLogFile sets the log file for the logger.
// The log file is used to write log messages.
// The fileName parameter specifies the name of the log file.
// If fileName is an empty string, the logger will not write log messages to a file.
// If fileName is provided, the logger will attempt to open or create the log file.
//   - If the file already exists, log messages will be appended to it.
//   - If the file doesn't exist, a new file will be created for logging.
//
// The log file should be closed appropriately after use to release system resources.
func (l *loggerOptions) SetLogFile(fileName string) {
	// open or create log file
	file := new(os.File)
	if fileName != "" {
		var err error
		file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file = nil
	}

	l.logFile = file
}
