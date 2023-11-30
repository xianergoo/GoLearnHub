package main

import "fmt"

// 判定从2到x-1都无法被它整除，就证明改数字是素数
func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	count := 0
	for i := 101; i <= 200; i++ {
		b := isPrime(i)
		if b {
			fmt.Printf("i: %v\n", i)
			count++
		}
	}
	fmt.Printf("count: %v\n", count)
}
