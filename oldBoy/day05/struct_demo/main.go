package main

import "fmt"

// 结构体是值类型
type person struct {
	name string
	age  int
}

func f(x person) {
	x.age = 20 // 修改的是副本的 age，不影响之前的值
}
func f2(x *person) {
	(*x).age = 20
}
func main() {
	var p person
	p.name = "张三"
	p.age = 19

	f(p)
	fmt.Println(p.age) //19
	f2(&p)
	fmt.Println(p.age) //20

}
