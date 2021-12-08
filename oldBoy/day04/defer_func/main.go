package main

import "fmt"

// Go 语言中函数不是原子操作，在底层分为两步
// 第一步： 返回值赋值
// 第二步： 真正的 RET 返回
// 如果存在 defer ,defer在第一步中间执行
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿") // defer 把后面的语句延迟到返回之前执行，用来关闭资源
	fmt.Println("end")
	defer fmt.Println("哈哈哈")
}

func main() {
	deferDemo()
}
