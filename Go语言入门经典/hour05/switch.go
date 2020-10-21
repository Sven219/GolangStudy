package main

import "fmt"

func main() {
	i := 33
	switch i {
	case 2:
		fmt.Println("TWO")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("I don't recognize that number")
	}
}
