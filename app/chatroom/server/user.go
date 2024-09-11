package main

import (
	"log"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	go user.ListenMessage()

	return user
}

func (u *User) Online() {
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()

	u.server.broadCast(u, "online")

}

func (u *User) Offline() {
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()

	u.server.broadCast(u, "offline")

}

func (u *User) SendMsg(msg string) {
	u.conn.Write([]byte(msg))
}

func (u *User) DoMessage(msg string) {
	log.Printf("user:%s: do message[%d]:%s\n", u.Name, len(msg), msg)

	if msg == "who" {
		//query online user
		var onlineMsg string
		u.server.mapLock.Lock()
		for _, user := range u.server.OnlineMap {
			onlineMsg = onlineMsg + "[" + user.Addr + "]" + user.Name + ": online\n"
		}
		u.server.mapLock.Unlock()
		if onlineMsg != "" {
			u.SendMsg(onlineMsg)
		} else {
			u.SendMsg("No User Online")
		}

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]

		if _, ok := u.server.OnlineMap[newName]; ok {
			u.SendMsg("this name is used")
		} else {
			u.server.mapLock.Lock()
			delete(u.server.OnlineMap, u.Name)
			u.server.OnlineMap[newName] = u
			u.server.mapLock.Unlock()

			u.Name = newName
			u.SendMsg("update name success")
		}

	} else if len(msg) > 4 && msg[:3] == "to|" {
		//msg format to|name|message

		//1 get username
		// get user
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			u.SendMsg("format incorrect, pls use \"to|:name|hello\" format like\n")
			return
		}
		remoteUser, ok := u.server.OnlineMap[remoteName]
		if !ok {
			u.SendMsg("user not exist.\n")
			return
		}

		//2 get message
		content := strings.Split(msg, "|")[2]
		// if content == "" {
		// }

		//3sendmessage
		remoteUser.SendMsg(u.Name + "says: " + content + "\n")
	} else {
		u.server.broadCast(u, msg)
	}
}

func (u *User) ListenMessage() {
	for {
		select {
		case msg, ok := <-u.C:
			if !ok {
				log.Printf("[%s]Listen Closed\n", u.Name)
				return
			}
			_, err := u.conn.Write([]byte(msg + "\n"))
			if err != nil {
				log.Println("write error", err)
				return
			}
		}
	}
}
