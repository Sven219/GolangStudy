package main

import (
	"fmt"
	"sync"
)

// channel

var a []int
var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)
	b = make(chan int) // 不带缓冲区的通道初始化
	// b = make(chan int, 16) // 带缓冲区的通道初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台 goroutine 从通道中取到了", x)
	}()
	fmt.Println("10 发送到通道 b 中了")
	b <- 10
	wg.Wait()
}
func bufChannel() {
	fmt.Println(b)
	// b = make(chan int) // 不带缓冲区的通道初始化
	b = make(chan int, 1) // 带缓冲区的通道初始化
	b <- 10
	fmt.Println("10 发送到通道 b 中了")
	x := <-b
	fmt.Println("后台 goroutine 从通道中取到了", x)
	b <- 20
	fmt.Println("20 发送到通道 b 中了")
	x1 := <-b
	fmt.Println("后台 goroutine 从通道中取到了", x1)
	close(b)
}

func main() {
	bufChannel()

}
