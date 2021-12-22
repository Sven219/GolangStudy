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

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now().Format("2016-01-02 15:04:05")
		fmt.Printf("[%s] [Debug] %s\n", now, msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now().Format("2016-01-02 15:04:05")
		fmt.Printf("[%s] [Info] %s\n", now, msg)
	}
}
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now().Format("2016-01-02 15:04:05")
		fmt.Printf("[%s] [Warning] %s\n", now, msg)
	}
}
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now().Format("2016-01-02 15:04:05")
		fmt.Printf("[%s] [Error] %s\n", now, msg)
	}
}
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now().Format("2016-01-02 15:04:05")
		fmt.Printf("[%s] [Fatal] %s\n", now, msg)
	}
}
