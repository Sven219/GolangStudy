package main

import (
	"fmt"
	"net"
)

// tcp server

func processConn(conn net.Conn) {
	// 3. 与客户端通信
	var tmp [128]byte

	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn failed ,err", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}

}

func main() {
	// 1.  本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	fmt.Println("Wait...")
	if err != nil {
		fmt.Println("Start server on 127.0.0.1:20000 failed：", err)
		return
	}
	// 1. 等待别人来链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed ,err :", err)
			return
		}
		go processConn(conn)

	}

}
