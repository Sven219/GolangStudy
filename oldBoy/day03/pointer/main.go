package main

import "fmt"

func main() {
	//1. & 取地址
	//2. * 根据地址取值

	n := 19
	fmt.Println(&n)
	p := &n
	fmt.Printf("%T\n", p)
	fmt.Println(*p)
}
