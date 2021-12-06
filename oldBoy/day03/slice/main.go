package main

import "fmt"

func main() {
	// 切片，拥有相同类型元素的可变长度
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 长度和容量
	fmt.Printf("len:%d cap:%d\n", len(s1), cap(s1))
	fmt.Printf("len:%d cap:%d\n", len(s2), cap(s2))
	// append 时容量翻倍
	s1 = append(s1, 3)
	s1 = append(s1, 4, 5, 7)
	fmt.Println(s1)
	fmt.Printf("len:%d cap:%d\n", len(s1), cap(s1))

	// 数组转切片
	a1 := [...]int{1, 3, 5, 7, 9}
	s3 := a1[0:2]
	fmt.Println(s3)
	fmt.Printf("len:%d cap:%d\n", len(s3), cap(s3))

}
