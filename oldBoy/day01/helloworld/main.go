package main

import (
	"fmt"
)

const pi = 3.14
const (
	a1 = iota
	a2
	a3
	a4
)

func main() {
	fmt.Println("hello world")
	data := "test"
	fmt.Printf("%s", data)
	fmt.Println()
	fmt.Print(pi)
	fmt.Println("aa1:", a1)
	fmt.Println("aa2", a2)
	// 修改字符串
	s1 := "白萝卜"
	s2 := []rune(s1)
	s2[0] = '红'
	fmt.Println(string(s2))
}
