package main

import "fmt"

func func_1() {
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	i := max(1, 2)
	fmt.Printf("i: %v\n", i)
}

func func_2() {
	func(a, b int) {
		max := 0
		if a > b {
			max = a
		} else {
			max = b
		}
		fmt.Printf("max: %v\n", max)
	}(1, 2)

}

func main() {
	func_1()
	func_2()
}
