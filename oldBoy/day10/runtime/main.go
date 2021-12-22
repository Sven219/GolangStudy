package main

import (
	"fmt"
	"path"
	"runtime"
)

func f1() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("error")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func main() {
	f1()
}
