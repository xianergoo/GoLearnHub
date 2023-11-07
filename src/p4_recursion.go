package main

import "fmt"

func func_resursion(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * func_resursion(n-1)
	}
}

func fibonacci_rec(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci_rec(n-1) + fibonacci_rec(n-2)
}

func main() {
	n := 5
	r := func_resursion(n)
	fmt.Printf("r: %v\n", r)

	x := 6
	for i := 0; i < x; i++ {
		fmt.Println(fibonacci_rec(i))
	}
}
