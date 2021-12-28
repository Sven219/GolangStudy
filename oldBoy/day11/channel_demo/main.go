package main

import (
	"fmt"
	"sync"
)

// channel 练习
// 1. 启动一个 goroutine，生成 100 个数发送到 ch1
// 2. 启动一个 goroutine，从 ch1 中取值，计算其平方放到 ch2
// 3. 在 main 函数中，从 ch2 取值打印出来

var wg sync.WaitGroup
var onece sync.Once

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	onece.Do(func() {
		close(ch2)
	})
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	for ret := range b {
		fmt.Println(ret)
	}
	// wg.Wait()
}
