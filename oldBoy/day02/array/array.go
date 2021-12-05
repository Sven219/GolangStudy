package main

import "fmt"

func main() {
	// 定义
	var a1 [3]bool
	var a2 [4]bool

	fmt.Printf("a1:%T  a2:%T\n", a1, a2)

	// 初始化
	// 如果不初始化，默认元素都是零值
	fmt.Println(a1, a2)
	// 初始化 1
	a1 = [3]bool{true, true, true}
	// 初始化 2
	a10 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a10)
	// 初始化 3
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	// 数组的遍历
	// 1. 索引遍历
	citys := [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	// 2. for range 遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}
	// 多维数组
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
	}
	fmt.Println(a11)

	// 多维数组遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
