package main

import (
	"fmt"
	"net"
	"strings"
)

// udp server

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("Listen UDP Error:", err)
		return
	}
	defer conn.Close()
	// 不需要建立链接，直接收发数据
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("Read from udp failed", err)
			return
		}
		fmt.Println(string(data[:n]))
		reply := strings.ToUpper(string(data[:n]))
		// 发送数据
		conn.WriteToUDP([]byte(reply), addr)
	}
}
