package main

import "fmt"

func main() {
	var l, m, z int = 1, 0, 0
	for i := 1; i < 13; i++ {
		fmt.Printf("month: %d, fresh: %d, middle: %d ,adult: %d, sum: %d\n", i, l, m, z, l+m+z)
		z += m
		m += l - m
		l = z
	}
}
