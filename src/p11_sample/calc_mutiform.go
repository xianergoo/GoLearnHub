package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d*%d=%-3d", i, j, i*j)
		}
		fmt.Println("")
	}
}
