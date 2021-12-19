package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFromFile() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed %v", err)
		return
	}
	// 关闭文件
	defer fileObj.Close()
	// 读取文件
	var tmp = make([]byte, 128)
	for {
		n, err := fileObj.Read(tmp)
		if err != nil {
			fmt.Printf("%v", err)
		}
		// fmt.Printf("读了%d 个字节\n", n)
		fmt.Println(string(tmp[:n]))

		if n < 128 {
			return
		}

	}
}

// bufio 读文件
func readFromFilebyBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed %v", err)
		return
	}
	// 关闭文件
	defer fileObj.Close()
	// 创建一个读的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		fmt.Println(line)
	}
}

// 打开文件
func main() {
	readFromFilebyBufio()

}
