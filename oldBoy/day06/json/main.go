package main

import (
	"encoding/json"
	"fmt"
)

// 结构体和 json

type person struct {
	Name string `json:"姓名",db:"name"`
	Age  int    `json:"年龄"`
}

func main() {
	p1 := person{
		Name: "张三",
		Age:  10,
	}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Marshal error")
		return
	}
	fmt.Printf("string(b): %v\n", string(b)) // string(b): {"姓名":"张三","年龄":10}

	// 反序列化
	str := "{\"姓名\":\"张三\",\"年龄\":10}"
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("p2: %v\n", p2)

}
