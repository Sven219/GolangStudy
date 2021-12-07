package main

import (
	"fmt"
)

func main() {
	var m1 map[string]int
	m1 = make(map[string]int, 10)
	m1["理想"] = 18
	m1["秀才"] = 19
	fmt.Println(m1)

	fmt.Println(m1["理想"])
	// 判断 map 是否含有某个值
	v, ok := m1["哪吒"]
	if !ok {
		fmt.Println("查无此人")

	} else {
		fmt.Println(v)
	}

	// 遍历

	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 只遍历 key
	for k := range m1 {
		fmt.Println(k)
	}
	// 只遍历 value
	for _, v := range m1 {
		fmt.Println(v)
	}
	// 删除
	delete(m1, "理想")
	fmt.Println(m1)

}
