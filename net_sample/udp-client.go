package main

import (
	"fmt"
	"net"
	"time"
)

var (
	serverAddr = "localhost:8080"
)

func main() {
	// 解析UDP地址
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 建立UDP连接
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 发送数据
	msg := []byte("Hello, server!")
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 接收响应
	buf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(time.Second * 5)) // 设置读取超时时间
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Received:", string(buf[:n]))
}
