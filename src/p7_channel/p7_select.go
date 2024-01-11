package main

import (
	"fmt"
	"time"
)

func test1(ch chan string) {
	fmt.Println("test1 begin")
	time.Sleep(time.Second * 5)
	ch <- "test1"
}

func test2(ch chan string) {
	fmt.Println("test2 begin")
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

func main() {
	// output1 := make(chan string)
	// output2 := make(chan string)

	// go test1(output1)
	// go test2(output2)

	// select {
	// case s1 := <-output1:
	// 	fmt.Println("s1=", s1)
	// case s2 := <-output2:
	// 	fmt.Println("s2=", s2)
	// }

	// time.Sleep(time.Second * 10)

	output1 := make(chan string, 10)
	go write(output1)
	for s := range output1 {
		fmt.Println("res: ", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
