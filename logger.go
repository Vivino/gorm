package gorm

import (
	"log"
	"os"
)

var (
	defaultLogger = Logger{log.New(os.Stdout, "\r\n", 0)}
	LogFormatter  func(values ...interface{}) (messages []interface{})
)

type logger interface {
	Print(v ...interface{})
}

// LogWriter log writer interface
type LogWriter interface {
	Println(v ...interface{})
}

// Logger default logger
type Logger struct {
	LogWriter
}

// Print format & print log
func (logger Logger) Print(values ...interface{}) {
	logger.Println(LogFormatter(values...)...)
}
