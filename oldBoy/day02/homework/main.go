package main

import "fmt"

func sum(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v
	}
	return sum
}
func findIndex(array []int, sum int) {
	for i, v1 := range array {
		for j := i + 1; j < len(array); j++ {
			if (v1 + array[j]) == sum {
				fmt.Printf("(%d,%d)", i, j)
			}
		}

	}

}

func main() {
	a1 := [...]int{1, 3, 5, 7, 8}
	fmt.Printf("数组之和为:%d\n", sum(a1[:]))
	findIndex(a1[:], 8)
}
