package main

import "fmt"

func sunNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sunNumbers(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Printf("The result is %v\n", result)
}
