package main

import (
	"fmt"
	"net"
)

var (
	serverAddr = "localhost:8080"
)

func main() {
	// 解析UDP地址
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	fmt.Println(addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 建立UDP连接
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Server started on", serverAddr)

	for {
		// 接收请求
		buf := make([]byte, 1024)
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Received %d bytes from %s: %s\n", n, clientAddr, string(buf[:n]))

		// 发送响应
		_, err = conn.WriteToUDP(buf[:n], clientAddr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Sent %d bytes to %s\n", n, clientAddr)
	}
}
