package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib2(number float64) float64 {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	return x
}

func fib(ch chan string, number float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)
	ch <- fmt.Sprintf("Fib(%.2f): %.2f\n", number, x)
}

var quit = make(chan bool)

func fib3(c chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}

func main() {
	start := time.Now()

	// size := 15
	// ch := make(chan string, size)
	// for i := 1; i < size; i++ {
	// 	go fib(ch, float64(i))
	// }

	// for i := 1; i < size; i++ {
	// 	n := fib2(float64(i))
	// 	fmt.Printf("Fib(%v): %v\n", i, n)
	// 	// fmt.Println(<-ch)
	// }

	// command := ""
	// data := make(chan int)
	// for {
	// 	num := <-data
	// 	fmt.Println(num)
	// 	fmt.Scanf("%s", &num)
	// 	if command == "quit" {
	// 		quit <- true
	// 		break
	// 	}
	// }\

	command := ""
	data := make(chan int)

	go fib3(data)

	for {
		num := <-data
		fmt.Println(num)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}

	time.Sleep(1 * time.Second)

	time.Sleep(1 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
