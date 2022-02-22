package main

import (
	"fmt"

	"../split_string"
)

func main() {
	ret := split_string.Split("babcbef", "b")
	fmt.Println(ret)
}
