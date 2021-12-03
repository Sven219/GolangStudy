package main

import (
	"fmt"
	"unicode"
)

func main() {
	var num int
	var f float32
	var b bool
	var str string

	num = 1
	f = 3.14
	b = true
	str = "test"

	fmt.Printf("类型：%T,值：%d  \n", num, num)
	fmt.Printf("类型：%T,值：%f  \n", f, f)
	fmt.Printf("类型：%T,值：%t  \n", b, b)
	fmt.Printf("类型：%T,值：%s  \n", str, str)

	str2 := "hello沙河小王子"
	str3 := []rune(str2)
	count := 0
	for i := 0; i < len(str3); i++ {
		if unicode.Is(unicode.Han, str3[i]) {
			count++
		}
	}
	fmt.Print(count)

}
