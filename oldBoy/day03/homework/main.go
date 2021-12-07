package main

import (
	"fmt"
	"strings"
)

// 统计一个字符串每个单词出现的次数

func countNum(str string) map[string]int {
	strarray := strings.Fields(strings.TrimSpace(str))
	m1 := make(map[string]int)
	for _, v := range strarray {
		m1[v]++
	}
	return m1
}

func main() {
	str := "What’s the differences between HashMap and Hashtable"
	fmt.Println(countNum(str))
}
