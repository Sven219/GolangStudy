package main

import "fmt"

func main() {
	// 拷贝
	a1 := []int{1, 2, 3}
	a2 := a1
	a3 := make([]int, len(a1))
	copy(a3, a1)
	fmt.Println(a1, a2, a3) // [1 2 3] [1 2 3] [1 2 3]
	a1[0] = 4
	fmt.Println(a1, a2, a3) // [4 2 3] [4 2 3] [1 2 3]

	// 删除元素
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 删除索引为 2 的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a)
}
