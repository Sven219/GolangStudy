package main

import "fmt"

// 递归：函数自己调用自己
// 递归一定要你有一个明确的退出条件

// 阶乘
func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

// n个台阶，一次可以走 1 步，也可以走 2 步，有多少种走法
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return taijie(n-1) + taijie(n-2)
	}
}

func main() {
	fmt.Println(fact(10))

	fmt.Println(taijie(3))
}
