package main

import "fmt"

func main() {
	// make 函数创建切片，指定长度和容量
	s1 := make([]int, 0, 5)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 1, 2, 3, 4, 5, 6)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))

	// 切片的赋值
	s3 := []int{1, 2, 3, 4, 6, 7, 8, 10}
	s4 := s3            // s3 和 s4 都指向了同一个底层数组
	fmt.Println(s3, s4) // [1 2 3 4 6 7 8 10] [1 2 3 4 6 7 8 10]
	s3[0] = 1000
	fmt.Println(s3, s4) //[1000 2 3 4 6 7 8 10] [1000 2 3 4 6 7 8 10]

	// 切片的遍历
	// 1. 索引遍历
	s5 := []int{1, 2, 3, 4, 6, 7, 8, 10}
	for i := 0; i < len(s5); i++ {
		fmt.Println(s5[i])
	}
	// 2. for range
	for i, v := range s5 {
		fmt.Println(i, v)
	}

}
