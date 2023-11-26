package main

import (
	"fmt"
)

//const constantName [type] = value

func main() {
	// const PI float32 = 3.14
	// const PI2 = 3.1415
	// fmt.Printf("PI: %v\n", PI)

	//iota
	const (
		a1 = iota
		_  = iota
		a2 = iota
	)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)

	/* 	const (
	   		a1 = iota
	   		a2 = iota
	   		a3 = iota
	   	)
	   	fmt.Println("a1: %v", a1)
	   	fmt.Println("a2: %v", a2)
	   	fmt.Println("a3: %v", a3) */
}
