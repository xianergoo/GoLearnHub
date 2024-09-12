package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	conn       net.Conn

	flag int //currnet clent mode
}

func NewClient(ip string, port int) (*Client, error) {
	client := &Client{
		ServerIP:   ip,
		ServerPort: port,
		flag:       -1,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("connected err: ", err)
		return nil, err
	}

	client.conn = conn

	return client, nil
}

func (c *Client) menu() bool {

	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.Exit")

	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		c.flag = flag
		return true
	} else {
		fmt.Println(">>please input valid number.")
		return false
	}
}

func (c *Client) updateName() bool {
	fmt.Println(">>>> please input new name")
	fmt.Scanln(&c.Name)

	sendMsg := "rename|" + c.Name + "\n"
	_, err := c.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn write err: ", err)
		return false
	}
	return true
}

func (c *Client) selectUser() {
	getUsersMesssage := "who"
	_, err := c.conn.Write([]byte(getUsersMesssage))
	if err != nil {
		fmt.Println("conn write err: ", err)
		return
	}

}

func (c *Client) publicChat() {
	fmt.Println(">>>> please input message, input \"exit\" exit chat")
	var chatmsg string
	fmt.Scanln(&chatmsg)
	for chatmsg != "exit" {
		if len(chatmsg) == 0 {
			fmt.Println(">>>> input message can not be empty, please input again")
			fmt.Scanln(&chatmsg)
			continue
		}
		_, err := c.conn.Write([]byte(chatmsg + "\n"))
		if err != nil {
			fmt.Println("conn write err: ", err)
			break
		}
		chatmsg = ""
		fmt.Println(">>>> please input message, input \"exit\" exit chat")
		fmt.Scanln(&chatmsg)
	}
}

func (c *Client) privateChat() {
	var remotName string
	for {
		c.selectUser()
		fmt.Println(">>>> please input username you want to chat with, input \"exit\" exit chat")
		fmt.Scanln(&remotName)
		if remotName == "" {
			continue
		}
		if remotName == "exit" {
			break
		}

		var chatMsg string
		for {
			if chatMsg == "exit" {
				break
			}
			fmt.Println(">>>> please input message, input \"exit\" exit chat")
			fmt.Scanln(&chatMsg)
			if chatMsg == "" {
				fmt.Println("message can not be empty, please input again")
				continue
			}
			formatMessage := fmt.Sprintf("to|%s|%s", remotName, chatMsg)
			_, err := c.conn.Write([]byte(formatMessage + "\n"))
			if err != nil {
				fmt.Println("conn write err: ", err)
				break
			}

		}

	}
}

func (c *Client) DealResponse() {
	//永久阻塞监听， 将conn的数据拷贝到stdout输出上面
	io.Copy(os.Stdout, c.conn)

	// for {
	// 	if c.conn == nil {
	// 		return
	// 	}
	// 	buf := make([]byte, 1024)
	// 	c.conn.Read(buf)
	// 	fmt.Println(buf)
	// }
}

func (c *Client) run() {
	for c.flag != 0 {
		for c.menu() != true {
			continue
		}

		switch c.flag {
		case 1:
			c.publicChat()
		case 2:
			c.privateChat()
		case 3:
			c.updateName()
		}
	}
}

var serverIP string
var serverPort int

func init() {
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "setting ip(default:127.0.0.1)")
	flag.IntVar(&serverPort, "port", 9001, "setting port(default9:001)")
}

func main() {
	flag.Parse()
	fmt.Println(serverIP, serverPort)
	client, err := NewClient(serverIP, serverPort)
	if err != nil {
		fmt.Println(">>>connected failed")
		return
	}
	fmt.Println(">>>>connected success")

	//start client
	// listen client
	go client.DealResponse()

	client.run()
}
