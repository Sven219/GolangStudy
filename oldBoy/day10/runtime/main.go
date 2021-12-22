package main

import (
	"fmt"
	"runtime"
)

func f1() {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("error")
		return
	}
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)
}

func main() {
	f1()
}
