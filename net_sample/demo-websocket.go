package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 将HTTP连接升级为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error: ", err)
		return
	}
	defer conn.Close()

	for {
		// 从WebSocket连接中读取消息
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error: ", err)
			break
		}

		// 打印接收到的消息
		log.Printf("recv: %s\n", message)

		// 将当前时间作为响应消息发送给客户端
		resp := []byte(time.Now().String())
		err = conn.WriteMessage(websocket.TextMessage, resp)
		if err != nil {
			log.Println("write error: ", err)
			break
		}
	}
}
