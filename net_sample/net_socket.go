package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		read := bufio.NewReader(conn)
		var buf [128]byte

		n, err := read.Read(buf[:])
		if err != nil {
			fmt.Println("read failed: err: ", err)
			break
		}

		revStr := string(buf[:n])
		fmt.Println("receive: ", revStr)
		conn.Write([]byte(revStr))
	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9001")
	if err != nil {
		fmt.Println("conn failed: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed: err: ", err)
			continue
		}

		go process(conn)
	}
}
