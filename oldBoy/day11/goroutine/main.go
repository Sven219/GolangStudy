package main

import (
	"fmt"
	"sync"
)

// goroutine

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个主 goroutine 去执行
func main() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		// go hello(i) // 开启一个单独的 goroutine 去执行 hello 函数
		go func() {
			fmt.Println(i)
		}()
	}
	fmt.Println("main")
}
