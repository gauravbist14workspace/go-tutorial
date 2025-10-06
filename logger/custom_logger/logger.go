package logger

import (
	"log"
	"os"
)

type LogLevel int

const (
	InfoLevel LogLevel = iota
	WarnLevel
	ErrorLevel
)

type Logger struct {
	level       LogLevel
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

var logger *Logger

func Init() {
	logger = &Logger{
		level:       InfoLevel,
		infoLogger:  log.New(os.Stdout, "INFO:", log.Lshortfile|log.LstdFlags),
		warnLogger:  log.New(os.Stdout, "WARN:", log.Ldate|log.LstdFlags),
		errorLogger: log.New(os.Stdout, "ERROR:", log.Ldate|log.Lshortfile),
	}
}

func SetLevel(level LogLevel) {
	logger.level = level
}

func Info(message string) {
	if logger != nil && logger.level <= InfoLevel {
		logger.infoLogger.Println(message)
	}
}

func Warn(message string) {
	if logger != nil && logger.level <= WarnLevel {
		logger.warnLogger.Println(message)
	}
}

func Error(message string) {
	if logger != nil && logger.level <= ErrorLevel {
		logger.errorLogger.Println(message)
	}
}
