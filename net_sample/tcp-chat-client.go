package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

var (
	serverAddr = "localhost:8080"
)

func main() {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalln(err)
		return
	}

	defer func() {
		log.Println("close self")
		_, err := conn.Write([]byte("ssss"))
		if err != nil {
			log.Println(err)
		}

		conn.Close()
	}()

	go func() {
		for {
			msg := make([]byte, 4096)
			n, err := conn.Read(msg)
			if err != nil {
				log.Fatalln(err)
			}

			log.Println(strings.TrimSpace(string(msg[:n])))

		}
	}()

	for {
		var msg string
		fmt.Scanln(&msg)
		if msg == "/quit" {
			return
		}
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Println(err)
			return
		}
	}

}
