package main

import (
	"fmt"
	"time"
)

func recv(c chan int) {
	ret := <-c
	fmt.Println("received: ", ret)
}

func main() {
	//无缓冲通道
	//1.
	ch := make(chan int)
	ch <- 10
	fmt.Println("send sucess")
	//build sucess, run error: fatal error: all goroutines are asleep - deadlock!

	//2.
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("send sucess")
	time.Sleep(time.Second)

	//有缓冲通道
	ch := make(chan int, 1)
	ch <- 10
	ch <- 11 //超过通道容量 runtime error: fatal error: all goroutines are asleep - deadlock!
	fmt.Println("send sucess")
}
