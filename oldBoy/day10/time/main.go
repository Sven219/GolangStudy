package main

import (
	"fmt"
	"time"
)

// 时间

func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())

	//时间戳

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	// time.Unix()
	ret := time.Unix(1640079446, 0)
	fmt.Println(ret)
	// 一个小时后的时间
	fmt.Println(now.Add(24 * time.Hour))

	// 定时器
	// timer := time.Tick(time.Second * 2)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	// 时间格式化
	fmt.Println(now.Format("2006年1月2日 15点4分5秒.000毫秒"))

	// 字符串时间转成时间戳
	t, err := time.Parse("2006-1-2 15点4分5秒", "1996-1-2 15点4分5秒")
	if err != nil {
		fmt.Printf("error:%s\n", err)
		return
	}
	fmt.Println(t.Unix())

}
