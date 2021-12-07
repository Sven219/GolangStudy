package main

import "fmt"

// 把一段逻辑抽象出来封装到一个函数中

// 有参数的函数有返回值
func f1(x int, y int) int {
	return x + y
}

// 没有参数没有返回值
func f2() {
	fmt.Println("无参无返回值")
}

// 没有参数有返回值
func f3() int {
	return 1
}

// 参数可以命名
func f4(x int, y int) (ret int) {
	ret = x + y
	return
}

// 多个返回值
func f5() (int, string) {
	return 1, "沙河"
}

// 参数的类型简写
func f6(x, y int) int {
	return x + y
}

// 可变长参数
func f7(x int, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}
func main() {

	fmt.Println(f4(1, 2))

	m, n := f5()
	fmt.Println(m, n)

	f7(1, 2, 3, 4, 5)
}
