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
	// 字符串转换成整型
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v,%T\n", retInt, retInt)
	i2 := 97
	ret4 := strconv.Itoa(i2)
	fmt.Printf("%#v,%T\n", ret4, ret4)
	// 从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v,%T\n", boolValue, boolValue)

}
