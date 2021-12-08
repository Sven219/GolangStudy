package main

import "fmt"

// 函数类型

func f1() {
	fmt.Println("Hello world")
}

func f2() int {
	fmt.Println("Hello world")
	return 2
}

func f3(x func() int) {
	ret := x()
	fmt.Println(ret)

}

func main() {
	a := f1
	fmt.Printf("%T\n", a)

	a2 := f2
	fmt.Printf("%T\n", a2)

	f3(f2)
}
