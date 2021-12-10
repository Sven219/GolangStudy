package main

import "fmt"

// 闭包
// 闭包是一个函数，这个函数能包含它外部作用域的一个变量，
// 底层原理
// 1. 函数可以作为返回值
// 2. 函数内部查找变量的顺序： 现先在自己内部找，找不到往外层找

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
