package main

import (
	"fmt"

	"github.com/gookit/ini"
)

func main() {
	// 打开INI文件
	cfg, err := ini.LoadFiles("config.ini")
	if err != nil {
		fmt.Println("无法打开INI文件:", err)
		return
	}

	// 读取键值对
	section, _ := cfg.Section("db")
	username := section.Key("username").String()
	password := section.Key("password").String()
	host := section.Key("host").String()
	port := section.Key("port").MustInt()

	fmt.Println("用户名:", username)
	fmt.Println("密码:", password)
	fmt.Println("主机:", host)
	fmt.Println("端口:", port)
}
