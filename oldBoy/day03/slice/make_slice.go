package main

import "fmt"

func main() {
	// make 函数创建切片，指定长度和容量
	s1 := make([]int, 0, 5)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 1, 2, 3, 4, 5, 6)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))

}
