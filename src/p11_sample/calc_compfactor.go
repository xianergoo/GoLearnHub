package main

import "fmt"

func main() {
	var m int
	for n := 2; n < 1000; n++ {
		m = n
		for i := 1; i < n; i++ {
			if n%i == 0 {
				m -= i
			}
		}
		if m == 0 {
			fmt.Printf("n: %v\n", n)
		}
	}
}
