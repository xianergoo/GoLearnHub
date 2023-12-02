package main

import "fmt"

func main() {
	number := 20
	a, b, sum := 1.0, 2.0, 0.0
	for i := 1; i <= number; i++ {
		sum += b / a
		a, b = b, a+b
	}
	fmt.Printf("sum: %v\n", sum)
}
