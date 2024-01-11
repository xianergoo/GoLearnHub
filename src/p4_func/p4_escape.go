package main

import "fmt"

type myStruct struct {
	x int
	y int
}

func createStruct() *myStruct {
	var ms myStruct
	return &ms
}

func main() {
	ms := createStruct()
	fmt.Println(ms)
}
