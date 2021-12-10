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

	fmt.Println(lihua.age)

}
