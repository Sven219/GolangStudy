package main

import (
	"time"

	"../mylogger"
)

// 测试自己写的日志库
func main() {
	log := mylogger.NewFileLogger("debug", "./", "go.log", 1024*10)
	for {
		log.Debug("这是一条 Debug 日志")
		log.Info("这是一条 Info 日志")
		id := 100
		name := "张三"
		log.Error("这是一条 Error 日志，id:%d ,name:%s", id, name)
		log.Warning("这是一条 warning  日志")
		log.Fatal("这是一条 Fatal 日志")
		time.Sleep(time.Second)
	}

}
