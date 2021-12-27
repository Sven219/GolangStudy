package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup

func f() {
	rand.Seed(time.Now().Unix()) // 保证每次执行的时候都不一样
	for i := 0; i < 5; i++ {
		r := rand.Intn(10) // 0<=x<10
		fmt.Println(r)
	}
}
func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main() {
	// f()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait() // 等待 wg 计数器减为 0
}
