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
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Get file info failed ,err:%v\n", err)
		return false
	}
	// 如果当前文件大小大于等于日志文件的最大值
	return fileInfo.Size() > f.maxFileSize
}
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割日志文件

	// 2. rename 日志文件 xx.log->xx.log.201908031709
	nowStr := time.Now().Format("2006010215040500")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("报错拉")
		return nil, err
	}

	logName := path.Join(f.filePath, fileInfo.Name()) // 拿到当前的日志文件完整路径
	newlogFileName := fmt.Sprintf("%s%s", logName, nowStr)
	os.Rename(logName, newlogFileName)
	// 1. 关闭当前的日志文件
	file.Close()
	// 3. 打开一个新的文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 将打开的新日志文件对象赋值给f.fileObj
	return fileObj, nil
}
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2016-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			fileObj, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = fileObj
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			// 如果要记录的日志是大于等于 ERROR 级别，我还要在 ERROR 日志文件中再记录一次
			if f.checkSize(f.errFileObj) {
				fileObj, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = fileObj
			}
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
