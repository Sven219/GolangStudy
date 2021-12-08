package main

import "fmt"

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿") // defer 把后面的语句延迟到返回之前执行，用来关闭资源
	fmt.Println("end")
	defer fmt.Println("哈哈哈")
}

func main() {
	deferDemo()
}
