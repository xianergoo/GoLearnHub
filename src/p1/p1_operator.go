package main

import "fmt"

func main() {

	/* //算数运算符
	a := 100
	b := 10

	fmt.Printf("(a + b): %v\n", (a + b))
	fmt.Printf("(a - b): %v\n", (a - b))
	fmt.Printf("(a * b): %v\n", (a * b))
	fmt.Printf("(a / b): %v\n", (a / b))
	fmt.Printf("(a %% b): %v\n", (a % b))

	a++
	fmt.Printf("a: %v\n", a)
	b--
	fmt.Printf("b: %v\n", b) */

	/*//关系运算符
	a := 1
	b := 2

	fmt.Printf("(a > b): %v\n", (a > b))
	fmt.Printf("(a < b): %v\n", (a < b))
	fmt.Printf("(a >= b): %v\n", (a >= b))
	fmt.Printf("(a <= b): %v\n", (a <= b))
	fmt.Printf("(a == b): %v\n", (a == b))
	fmt.Printf("(a != b): %v\n", (a != b)) */

	/* //逻辑运算符
	a := true
	b := false

	fmt.Printf("(a && b): %v\n", (a && b))
	fmt.Printf("(a || b): %v\n", (a || b))
	fmt.Printf("(!a): %v\n", (!a))
	fmt.Printf("(!b): %v\n", (!b)) */

	/* //位运算符

	a := 4 // 二进制 100
	fmt.Printf("a: %b\n", a)
	b := 8 // 二进制 1000
	fmt.Printf("b: %b\n", b)

	fmt.Printf("(a & b): %v, %b \n", (a & b), (a & b))
	fmt.Printf("(a | b): %v, %b\n", (a | b), (a | b))
	fmt.Printf("(a ^ b): %v, %b\n", (a ^ b), (a ^ b))
	fmt.Printf("(a << 2): %v, %b\n", (a << 2), (a << 2))
	fmt.Printf("(b >> 2): %v, %b\n", (b >> 2), (b >> 2)) */

	//赋值运算符

	var a int
	a = 100
	fmt.Printf("a: %v\n", a)
	a += 1 // a = a + 1
	fmt.Printf("a: %v\n", a)
	a -= 1 // a = a -1
	fmt.Printf("a: %v\n", a)
	a *= 2 // a = a * 2
	fmt.Printf("a: %v\n", a)
	a /= 2 // a = a / 2
	fmt.Printf("a: %v\n", a)

}
