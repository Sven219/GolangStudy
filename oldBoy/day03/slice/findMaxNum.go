package main

import (
	"fmt"
	"sort"
)

func findNum(array []int) {
	sort.Ints(array)
	fmt.Printf("最大的数：%d\n", array[len(array)-1])
	for i := len(array) - 1; i > 0; i-- {
		if array[i] == array[i-1] {
			continue
		} else {
			fmt.Printf("第二大的数：%d\n", array[i-1])
			break
		}
	}
	return

}

func main() {
	nums := []int{22, 3343, 565, 3343, 332, 122, 444, 122}
	findNum(nums)
}
