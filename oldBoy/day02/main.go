package main

import (
	"fmt"
)

func main() {
	// 当 i=5 时跳出 for 循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("-----------")
	// 当 i=5 时 ，跳出此次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	// switch
	n := 3
	switch n {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("yes")
	}

}
