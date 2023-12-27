// package main

// import (
// 	"log"
// 	"os"
// )

// func init() {
// 	// 配置日志输出格式
// 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
// 	// 配置前缀
// 	log.SetPrefix("order:")
// 	// 配置输出位置
// 	logFile, err := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
// 	if err != nil {
// 		log.Panic("打开日志文件异常")
// 	}
// 	log.SetOutput(logFile)
// }

// func main() {
// 	log.Println("golang技术栈-老郭，log标准库测试")

// 	logger := log.Default()
// 	logger.SetFlags(log.Llongfile)
// 	logger.Output(0, "0 calldepth")
// 	logger.Output(1, "1 calldepth")
// 	logger.Output(2, "2 calldepth")
// }

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"log"
// )

// func TestLogger() {
// 	var (
// 		buf    bytes.Buffer
// 		logger = log.New(&buf, "logger: ", log.Lshortfile)
// 	)
// 	logger.Print("golang技术栈，go测试!")
// 	fmt.Print(&buf)
// }

// func main() {
// 	TestLogger()
// }

package main

import (
	"fmt"
	"log"
)

func TestLogPrint() {
	log.Println("TestLogPrint")
}

func TestLogPanic() {
	defer fmt.Println("Panic defer")
	log.Panicln("TestLogPanic")
}

func TestLogFatal() {
	defer fmt.Println("Panic Fatal")
	log.Fatalln("TestLogFatal")
}

func main() {
	TestLogPrint()
	TestLogPanic()
	// TestLogFatal()
}
