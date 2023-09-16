package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type Logger interface {
	Fatal(message string)
	Fatalf(format string, a ...interface{})
	Panic(message string)
	Panicf(format string, a ...interface{})
	Critical(message string)
	Criticalf(format string, a ...interface{})
	Error(message string)
	Errorf(format string, a ...interface{})
	Warning(message string)
	Warningf(format string, a ...interface{})
	Notice(message string)
	Noticef(format string, a ...interface{})
	Info(message string)
	Infof(format string, a ...interface{})
	Debug(message string)
	Debugf(format string, a ...interface{})
}

type logger struct {
	log      *log.Logger
	logLevel LogLevel
}

func NewLogger(logLevel LogLevel) Logger {
	return &logger{
		log:      log.New(os.Stdout, "", 0),
		logLevel: logLevel,
	}
}

func (l *logger) Fatal(message string) {
	l.printf(CriticalLevel, message)
	os.Exit(1)
}

func (l *logger) Fatalf(format string, a ...interface{}) {
	l.printf(CriticalLevel, format, a...)
	os.Exit(1)
}

func (l *logger) Panic(message string) {
	l.printf(CriticalLevel, message)
	panic(message)
}

func (l *logger) Panicf(format string, a ...interface{}) {
	l.printf(CriticalLevel, format, a...)
	message := fmt.Sprintf(format, a...)
	panic(message)
}

func (l *logger) Critical(message string) {
	l.printf(CriticalLevel, message)
}

func (l *logger) Criticalf(format string, a ...interface{}) {
	l.printf(CriticalLevel, format, a...)
}

func (l *logger) Error(message string) {
	l.printf(ErrorLevel, message)
}

func (l *logger) Errorf(format string, a ...interface{}) {
	l.printf(ErrorLevel, format, a...)
}

func (l *logger) Warning(message string) {
	l.printf(WarningLevel, message)
}

func (l *logger) Warningf(format string, a ...interface{}) {
	l.printf(WarningLevel, format, a...)
}

func (l *logger) Notice(message string) {
	l.printf(NoticeLevel, message)
}

func (l *logger) Noticef(format string, a ...interface{}) {
	l.printf(NoticeLevel, format, a...)
}

func (l *logger) Info(message string) {
	l.printf(InfoLevel, message)
}

func (l *logger) Infof(format string, a ...interface{}) {
	l.printf(InfoLevel, format, a...)
}

func (l *logger) Debug(message string) {
	l.printf(DebugLevel, message)
}

func (l *logger) Debugf(format string, a ...interface{}) {
	l.printf(DebugLevel, format, a...)
}

func (l *logger) printf(logLevel LogLevel, format string, a ...interface{}) {
	if l.logLevel < logLevel {
		return
	}
	text := fmt.Sprintf(format, a...)
	timeText := time.Now().Format(timeFormat)
	l.log.Printf("%v %v %v ▶ %s\033[0m", logLevel.Color().colorString(), timeText, logLevel.String(), text)
}
