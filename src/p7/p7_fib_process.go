package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(ch chan float64, number float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	ch <- fmt.Sprintf("Fib(%v): %v\n", number, x)
}

func main() {
	start := time.Now()

	size := 15
	ch := make(chan float64, size)
	for i := 1; i < size; i++ {
		go fib(ch, float64(i))
	}
	for i := 1; i < size; i++ {
		// fmt.Printf("Fib(%v): %v\n", i, n)
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
