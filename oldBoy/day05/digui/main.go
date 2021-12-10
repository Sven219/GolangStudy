package main

import "fmt"

// 递归：函数自己调用自己
// 递归一定要你有一个明确的退出条件

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	fmt.Println(fact(10))
}
