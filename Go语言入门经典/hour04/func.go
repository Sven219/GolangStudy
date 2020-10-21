package main

import "fmt"

func isEven(i int) bool {
	return i%2 == 0
}

func main() {
	fmt.Println(isEven(1))
	fmt.Println(isEven(2))
}
