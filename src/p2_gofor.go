package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	simple_for()
	f1()
}

func simple_for() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += 1
	}
	fmt.Println("sum of 1..100 is ", sum)
}

func f1() {
	var num int32
	sec := time.Now().Unix()
	rand.Seed(sec)

	for {
		fmt.Print("Writing inside the loop...")
		if num = rand.Int31n(10); num == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(num)
	}
}
