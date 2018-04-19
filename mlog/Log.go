package mlog

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// DebugEnabled controls whether Debug log messages are generated
var DebugEnabled = false

// InfoEnabled controls whether "regular" log messages are generated
var InfoEnabled = true

// ErrorEnabled controls whether Error log messages are generated
var ErrorEnabled = true

var debugLogger, infoLogger, errorLogger *log.Logger

func init() {
	debugLogger = log.New(os.Stderr, "[DEBUG] ", log.Lshortfile)
	infoLogger = log.New(os.Stderr, "[LOG] ", log.Lshortfile)
	errorLogger = log.New(os.Stderr, "[ERROR] ", 0)
}

// Debugf writes a debug message to stderr
func Debugf(format string, v ...interface{}) {
	if DebugEnabled {
		s := fmt.Sprintf(format, v...)
		_ = debugLogger.Output(2, s)
	}
}

// Printf writes a regular log message to stderr
func Printf(format string, v ...interface{}) {
	if InfoEnabled {
		s := fmt.Sprintf(format, v...)
		_ = infoLogger.Output(2, s)
	}
}

// Printv writes a variable as a regular log message to stderr
func Printv(v interface{}) {
	if InfoEnabled {
		//s := fmt.Sprintf("%#v", v)
		b, err := json.MarshalIndent(v, "", "    ")
		if err != nil {
			panic(err)
		}
		s := string(b)
		_ = infoLogger.Output(2, s)
	}
}

// Error writes an error message to stderr
func Error(err error) {
	if ErrorEnabled {
		s := err.Error()
		_ = errorLogger.Output(2, s)
	}
}
