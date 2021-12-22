package main

import (
	"fmt"
	"time"
)

func f1() {

}

// 时区
func f2() {
	now := time.Now()
	fmt.Println(now)
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2021-12-20 13:12:11", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	// 时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)

}

func main() {
	f2()
}
