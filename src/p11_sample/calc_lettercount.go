package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i, j, k, l = 0, 0, 0, 0
	fmt.Print("请输入一串字符：")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	for _, ch := range input {
		switch {
		case ch >= 'A' && ch <= 'Z':
			i++
		case ch >= 'a' && ch <= 'z':
			i++
		case ch == ' ' || ch == '\t':
			j++
		case ch >= '0' && ch <= '9':
			k++
		default:
			l++
		}
	}
	fmt.Printf("char count = %d，space count = %d，digit count = %d，others count =%d\n", i, j, k, l)
}
