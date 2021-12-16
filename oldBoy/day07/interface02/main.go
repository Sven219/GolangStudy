package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫动")
}
func (c cat) eat(s string) {
	fmt.Printf("吃%s", s)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("激动！")
}
func (c chicken) eat() {
	fmt.Println("吃鸡饲料")
}

func main() {
	var a1 animal

	bc := cat{
		name: "淘气",
		feet: 4,
	}

	a1 = bc
	a1.eat("鱼")
	fmt.Println(a1)

}
