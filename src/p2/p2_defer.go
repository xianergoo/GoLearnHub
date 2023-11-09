package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()
	// func_defer()
	// sample_defer()
	highlow(2, 0)
	fmt.Println("Program finished")
}

// use defer down will print regular list, and print -4, -3, -2, -1 desc
func func_defer() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("defferred", -i) //list append(-1, -2, -3, -4) 调用输出顺序后进先出
		fmt.Println("regular", i)          //print list 1, 2, 3, 4
	}
}

// defer 函数的一个典型用例是在使用完文件后将其关闭。 下面是一个示例
func sample_defer() {
	newfile, error := os.Create("learn.txt")
	if error != nil {
		fmt.Println("Error: counld not create file.")
		return
	}
	defer newfile.Close()

	if _, error = io.WriteString(newfile, "Learning GO!"); error != nil {
		fmt.Println("Error: could not write to file")
		return
	}
	newfile.Sync()
}

// panic
func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic")
		panic("highlow() low greater then hign")
	}
	defer fmt.Println("deffered: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}
