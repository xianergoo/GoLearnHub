package main

import "net"

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

	} else {
		u.server.broadCast(u, msg)
	}
}

func (u *User) ListenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
