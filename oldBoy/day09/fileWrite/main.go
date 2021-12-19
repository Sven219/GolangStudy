package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容

func writeDemo1() {
	fileObj, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	// write
	fileObj.Write([]byte("hello world"))
	// write string
	fileObj.WriteString("\n你好世界")
	defer fileObj.Close()
}
func writeDemo2() {
	fileObj, err := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer fileObj.Close()
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("你好啊，树哥")
	wr.Flush()
}

func writeDemo3() {
	str := "hello shuge"
	err := ioutil.WriteFile("./tes.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("error")
		return
	}

}

func main() {
	// writeDemo1()
	// writeDemo2()
	writeDemo3()

}
