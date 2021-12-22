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
func log(lv LogLevel, msg string) {
	now := time.Now().Format("2016-01-02 15:04:05")
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)

}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}
}
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		log(WARNING, msg)

	}
}
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		log(FATAL, msg)
	}
}
