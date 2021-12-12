package main

import "fmt"

type dog struct {
	name string
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
func (d dog) wang() {
	fmt.Println("%s:旺旺旺~", d.name)
}

func main() {
	d1 := newDog("旺财")
	d1.wang()
}
