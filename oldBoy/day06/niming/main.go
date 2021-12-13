package main

import "fmt"

// 匿名字段
type person struct {
	string
	int
}

func main() {
	p1 := person{
		"张三",
		10,
	}
	fmt.Println(p1.int)
	fmt.Println(p1.string)

}
