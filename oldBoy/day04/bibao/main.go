package main

import "fmt"

// 闭包

// func f1(f func()) {
// 	fmt.Println("This is f1")

// }
func f2(x, y int) {
	fmt.Println("This is f2")
	fmt.Println(x + y)
}

// 定义一个函数对 f2 进行包装
func f3(f func(int, int), m int, n int) func() {
	tmp := func() {
		f(m, n)
	}
	return tmp
}
func main() {
	ret := f3(f2, 100, 200)
	ret()
}
