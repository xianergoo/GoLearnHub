package main

import (
	"fmt"
	"os"
)

// 打开关闭文件
func openCloseFile() {
	// 只能读
	f, err := os.Open("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}

	// 根据第二个参数 可以读写或者创建
	f2, _ := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Printf("f2.Name(): %v\n", f2.Name())

	err := f.Close()
	fmt.Printf("err: %v\n", err)
	err2 := f2.Close()
	fmt.Printf("err2: %v\n", err2)
}

// 创建文件
func createFile() {
	// 等价于：OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	f, _ := os.Create("a2.txt")
	fmt.Printf("f.Name(): %v\n", f.Name())
	// 第一个参数 目录默认：Temp 第二个参数 文件名前缀
	f2, _ := os.CreateTemp("", "temp")
	fmt.Printf("f2.Name(): %v\n", f2.Name())
}

// 读操作
func readOps() {
	// 循环读取
	/* 	f, _ := os.Open("a.txt")
	   	for {
	   		buf := make([]byte, 6)
	   		n, err := f.Read(buf)
	   		fmt.Println(string(buf))
	   		fmt.Printf("n: %v\n", n)
	   		if err == io.EOF {
	   			break
	   		}
	   	}
	   	f.Close()
	*/
	/* buf := make([]byte, 10)
	f2, _ := os.Open("a.txt")
	// 从5开始读10个字节
	n, _ := f2.ReadAt(buf, 5)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f2.Close() */

	// 测试 a目录下面有b和c目录
	/* f, _ := os.Open("a")
	de, _ := f.ReadDir(-1)
	for _, v := range de {
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
		fmt.Printf("v.Name(): %v\n", v.Name())
	} */

	// 定位
	f, _ := os.Open("a.txt")
	f.Seek(3, 0)
	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f.Close()

}

func write() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_TRUNC, 0755)
	f.Write([]byte(" hello golang"))
	f.Close()
}

func writeString() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0755)
	f.WriteString("hello java")
	f.Close()
}

func writeAt() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0755)
	f.WriteAt([]byte("aaa"), 3)
	f.Close()

}

func main() {
	// openCloseFile()
	// createFile()
	// readOps()

	// write()
	// writeString()
	// writeAt()
}
