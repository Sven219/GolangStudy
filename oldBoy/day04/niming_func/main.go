package main

import "fmt"

func main() {
	f1 := func(x, y int) int {
		return x + y
	}

	fmt.Println(f1(10, 20))
}
