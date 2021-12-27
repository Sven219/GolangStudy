package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {
	i := int32(2317)
	ret := string(i)
	fmt.Println(ret)
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v", ret2)

	str := "10000"
	ret3, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return
	}
	fmt.Printf("%#v,%T\n", ret3, ret3)

}
