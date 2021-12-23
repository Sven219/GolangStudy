package mylogger

import (
	"fmt"
	"time"
)

// 自定义一个日志库

// Logger 日志对象
type Logger struct {
	Level LogLevel
}

// NewLog 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) enable(logLevel LogLevel) bool {
	return l.Level <= logLevel
}
func (l Logger) log(lv LogLevel, format string, a ...interface{}) {
	if l.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2016-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
	}
}
func (l Logger) Debug(format string, a ...interface{}) {
	l.log(DEBUG, format, a)
}
func (l Logger) Info(format string, a ...interface{}) {
	l.log(INFO, format, a)
}
func (l Logger) Warning(format string, a ...interface{}) {
	l.log(WARNING, format, a)
}
func (l Logger) Error(format string, a ...interface{}) {
	l.log(ERROR, format, a)
}
func (l Logger) Fatal(format string, a ...interface{}) {
	l.log(FATAL, format, a)
}
