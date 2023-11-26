package main

import "fmt"

func swap(a, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

func swap_(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	a := 10
	b := 20
	swap_(&a, &b)
	fmt.Printf("a: %v, b: %v\n", a, b)

	var p *int
	p = &a
	fmt.Printf("&a: %v\n", &a)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("*p : %v\n", *p)

	var pp **int
	pp = &p
	fmt.Printf("p: %v\n", &p)
	fmt.Printf("pp: %v\n", pp)
	fmt.Printf("pp: %v\n", *pp)
}
