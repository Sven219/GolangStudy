package main

import "fmt"

// 空接口
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

// interface 关键字
// interface{} 空接口
func main() {

	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "张三"
	m1["age"] = 19
	m1["married"] = true
	m1["hobby"] = [...]string{"唱", "跳"}
	fmt.Println(m1)

	show(false)
	show(nil)
	show(m1)
}
