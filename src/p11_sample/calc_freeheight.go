package main

import "fmt"

func main() {
	n := 10
	s := 100.0
	h := s / 2
	for i := 2; i <= n; i++ {
		s += h * 2
		h /= 2
	}
	fmt.Printf("s: %v\n", s)
	fmt.Printf("h: %v\n", h)
}
