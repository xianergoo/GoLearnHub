package main

import "fmt"

func fibonacci(n int) []int {

	if n < 2 {
		return make([]int, 0)
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if i < 2 {
			res[i] = 1
		} else {
			res[i] = res[i-1] + res[i-2]
		}

	}
	return res
}

func main() {
	x := 6
	s := fibonacci(x)
	fmt.Printf("s: %v\n", s)
	// fmt.Printf("s: %v\n", s)

}
