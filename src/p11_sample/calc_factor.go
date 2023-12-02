package main

import "fmt"

func fact(n int) int {
	var sum = 0
	if n == 0 {
		return 1
	}
	sum = n * fact(n-1)
	return sum

}
func main() {
	fmt.Printf("fact(5): %v\n", fact(5))
}
