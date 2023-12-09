package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("error connecting: %v\n", err)
		return
	}

	defer listener.Close()

	// fmt.Fprintf(listen, "server / HTTP/1.0\r\n\r\n")

	buf := make([]byte, 1024)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		n, _ := conn.Read(buf)
		if n == 0 {
			continue
		}
		fmt.Println(string(buf[0:n]))
	}
}
