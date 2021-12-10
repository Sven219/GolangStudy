package main

import "fmt"

// 结构体
type person struct {
	name  string
	age   int
	hobby []string
}

func main() {

	var lihua person
	lihua.age = 19
	lihua.hobby = []string{"吃", "喝"}
	lihua.name = "李华"
	fmt.Printf("%T\n", lihua)
	fmt.Println(lihua)

	// 匿名结构体
	var s struct {
		name string
		age  int
	}
	s.age = 1
	s.name = "张三"

	fmt.Printf("%T\n", s) //struct { name string; age int }
	fmt.Println(s) // {张三 1}
}
