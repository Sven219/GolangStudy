package main

import (
	"fmt"
	"time"
)

// 时间

func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
}
