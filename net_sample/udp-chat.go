package main

import (
	"fmt"
	"net"
	"strings"
)

const (
	serverAddr = "localhost:8080"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Server Started on", serverAddr, addr)

	clients := make(map[string]*net.UDPAddr)

	for {
		buf := make([]byte, 2048)
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}

		msg := strings.TrimSpace(string(buf[:n]))
		if msg == "" {
			continue
		}
		fmt.Printf("Received %d bytes from %s: %s\n", n, clientAddr, msg)
		if _, ok := clients[clientAddr.String()]; !ok {
			fmt.Println("New client:", clientAddr)
			clients[clientAddr.String()] = clientAddr
		}

		for _, addr := range clients {
			if addr.String() == clientAddr.String() {
				continue
			}

			_, err := conn.WriteToUDP([]byte(msg), addr)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Sent %d bytes to %s: %s\n", len(msg), addr, msg)
		}
	}
}
