package main

import (
	"log"
	"net"
)

var (
	clients    = make(map[net.Addr]net.Conn)
	addCh      = make(chan net.Conn)
	delCh      = make(chan net.Addr)
	messageCh  = make(chan []byte)
	listenAddr = "localhost:8080"
)

func broadcaster() {
	for {
		select {
		case conn := <-addCh:
			clients[conn.RemoteAddr()] = conn
			log.Println("new Client: :", conn.RemoteAddr())
		case addr := <-delCh:
			delete(clients, addr)
			log.Println("close Client: :", addr)
		case msg := <-messageCh:
			for _, conn := range clients {
				_, err := conn.Write(msg)
				if err != nil {
					log.Println(err)
				}

			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		delCh <- conn.RemoteAddr()
		conn.Close()
	}()

	for {
		msg := make([]byte, 4096)
		n, err := conn.Read(msg)
		if err != nil {
			return
		}
		messageCh <- msg[:n]
	}
}

func main() {
	log.Println("Server Started", listenAddr)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer listener.Close()

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		addCh <- conn

		go handleConn(conn)
	}

}
