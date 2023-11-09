package main

import "fmt"

var a int = initVar()

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func initVar() int {
	fmt.Println("init var ...")
	return 100
}

// init先于main执行
func main() {
	fmt.Println("main")
}
