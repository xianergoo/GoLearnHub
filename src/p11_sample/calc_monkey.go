package main

import (
	"fmt"
)

func main() {
	var x, y, day = 1, 0, 9

	for day > 0 {
		y = (x + 1) * 2
		x = y
		day--
	}

	fmt.Printf("一共有桃子： %d\n", y)
}
