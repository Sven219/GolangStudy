package main

import "fmt"

// 结构体模拟实现其它语言的继承

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s 会动！\n", a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d dog) wang() {
	fmt.Printf("%s 在叫旺旺旺\n", d.name)
}

func main() {
	d1 := dog{
		feet:   4,
		animal: animal{"旺财"},
	}
	d1.wang()
	d1.move()
}
