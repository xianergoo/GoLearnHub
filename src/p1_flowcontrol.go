package main

import (
	"fmt"
)

func main() {
	fmt.Println("aaa")
	fmt.Println("bbb")

	a := 100
	if a > 20 {
		fmt.Println("aaa")
	} else {
		fmt.Println("bbb")
	}

	switch a {
	case 100:
		fmt.Println("aaa")
	default:
		fmt.Println("bbb")

	}

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}

	x := [...]int{1, 2, 2}
	for _, v := range x {
		fmt.Printf("v: %v\n", v)
	}

}
