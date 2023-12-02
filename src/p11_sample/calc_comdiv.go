package main

import (
	"fmt"
)

func main() {
	var i1, i2, m, n, r, x int
	fmt.Println("input two nums: ")
	fmt.Scanf("%d%d", &i1, &i2)
	if i1 > i2 {
		m = i1
		n = i2
	} else {
		m = i2
		n = i1
	}

	x = m * n
	r = m % n
	for r != 0 {
		m = n
		n = r
		r = m % n
	}
	fmt.Printf("max common div: %d, min common muti: %d\n", n, x/n)

}
