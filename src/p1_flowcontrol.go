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

	//初始变量可以声明在布尔表达式中
	if age := 16; age > 18 {
		fmt.Println("you are adult")
		fmt.Printf("age: %v\n", age)
	} else if age < 12 {
		fmt.Println("you are child")
		fmt.Printf("age: %v\n", age)
	} else {
		fmt.Println("you are teen")
		fmt.Printf("age: %v\n", age)
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

	//

}
