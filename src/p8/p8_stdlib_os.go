package main

import (
	"fmt"
	"os"
)

// 创建文件
func createFile() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f: %v\n", f)
	}
}

// 创建目录
func createDir() {
	// 创建单个目录
	/* err := os.Mkdir("test", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} */
	err := os.MkdirAll("test/a/b", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 删除目录
func removeDir() {
	/* err := os.Remove("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} */

	err := os.RemoveAll("test")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 获得工作目录
func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}
}

// 修改工作目录
func chWd() {
	err := os.Chdir("d:/")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(os.Getwd())
}

// 获得临时目录
func getTemp() {
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

// 重命名文件
func renameFile() {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 读文件
func readFile() {
	b, err := os.ReadFile("test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("b: %v\n", string(b[:]))
	}
}

// 写文件
func writeFile() {
	s := "hello world"
	os.WriteFile("test3.txt", []byte(s), os.ModePerm)
}

func main() {
	// createFile()
	// createDir()
	// removeDir()
	// removeDir()
	// getWd()
	// chWd()
	// renameFile()
	// readFile()
	writeFile()
	getTemp()
}
