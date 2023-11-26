package main

import (
	"fmt"
	"log"
	"os"
)

// func Flags() int  // 返回标准log输出配置
// func SetFlags(flag int)  // 设置标准log输出配置

// const (
//     // 控制输出日志信息的细节，不能控制输出的顺序和格式。
//     // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
//     Ldate         = 1 << iota     // 日期：2009/01/23
//     Ltime                         // 时间：01:23:23
//     Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
//     Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
//     Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
//     LUTC                          // 使用UTC时间
//     LstdFlags     = Ldate | Ltime // 标准logger的初始值
// )

// func Prefix() string  // 返回日志的前缀配置
// func SetPrefix(prefix string)  // 设置日志前缀

func main() {

	// defer fmt.Println("发生了 panic错误！")
	// log.Print("my log")
	// log.Printf("my log %d", 100)
	// name := "tom"
	// age := 20
	// log.Println(name, ",", age)
	// log.Panic("致命错误！")
	// // log.Fatal("致命错误！")
	// fmt.Println("end...")
	i := log.Flags()
	fmt.Printf("i: %v\n", i)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("log test")

	s := log.Prefix()
	fmt.Printf("s: %v\n", s)
	log.SetPrefix("MyLog: ")
	log.Print("test log")
	s = log.Prefix()
	fmt.Println("s: %v\n", s)

	f, err := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic(err)
	}
	log.SetOutput(f)
	log.Print("my log...")
	logger.Println("自定义logger")
}

var logger *log.Logger

func init() {
	fmt.Println("init begin")
	logFile, err := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}

	logger = log.New(logFile, "success", log.Ldate|log.Ltime|log.Lshortfile)
	fmt.Println("init end")
}
