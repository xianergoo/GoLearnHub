package main

import (
	"fmt"
	"time"
)

func send(ch chan string, message string) {
	ch <- message
}

func send2(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
	// <-ch
}

func read2(ch <-chan string) {
	fmt.Printf("Receiveing: %#v\n", <-ch)
	// ch <- "hello" //cannot send to receive-only channel ch (variable of type <-chan string)
}

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing!"
}
func replicae(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating!"
}

func main() {
	ch3 := make(chan string)
	ch4 := make(chan string)
	go process(ch3)
	go replicae(ch4)

	ch2 := make(chan string, 1)
	send2(ch2, "hello")
	read2(ch2)

	size := 2
	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	go send(ch, "three")
	go send(ch, "four")
	fmt.Println("All data send to the channel...")

	for i := 0; i < 2; i++ {
		select {
		case process := <-ch3:
			fmt.Println(process)
		case replicae := <-ch4:
			fmt.Println(replicae)
		}
	}

	// ch2 := make(chan string, 1)
	// send2(ch2, "hello")
	// read2(ch2)
	// size := 2
	// ch := make(chan string, size)
	// send(ch, "one")
	// send(ch, "two")
	// go send(ch, "three")
	// go send(ch, "four")
	// fmt.Println("All data send to the channel...")

	for i := 0; i < size; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done!")
}
