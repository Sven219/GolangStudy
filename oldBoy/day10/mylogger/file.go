package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志相关的代码

type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	maxFileSize int64
	fileObj     *os.File
	errFileObj  *os.File
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       LogLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() // 按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() error {
	fullFilePath := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open log file failed,err:%s", err)
		return err
	}
	errfileObj, err := os.OpenFile(fullFilePath+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Open err log file failed,err:%s", err)
		return err
	}
	// 日志文件都已经打开
	f.fileObj = fileObj
	f.errFileObj = errfileObj
	return nil

}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2016-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			// 如果要记录的日志是大于等于 ERROR 级别，我还要在 ERROR 日志文件中再记录一次
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

func (f FileLogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

func (f FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a)
}

func (f FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a)
}
func (f FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a)
}
func (f FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a)
}
func (f FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a)
}
