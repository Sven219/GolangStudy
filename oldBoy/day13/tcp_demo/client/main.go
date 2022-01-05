package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//tcp  client

func main() {
	// 1. 与 server 端建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Dial failed ,err", err)
		return
	}
	// 2. 发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}

	conn.Close()
}
