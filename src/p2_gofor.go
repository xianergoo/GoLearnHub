package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	simple_for()
	f1()
	func_continue()
}

func simple_for() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += 1
	}
	fmt.Println("sum of 1..100 is ", sum)
}

// 无限循环和 break 语句
func f1() {
	// var num int32
	sec := time.Now().Unix()
	rand.Seed(sec)
	var num int
	for {
		fmt.Print("Writing inside the loop...")
		if num = rand.Intn(10); num == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(num)
	}
}

// Continue 语句
func func_continue() {
	sum := 0
	for num := 0; num <= 100; num++ {
		if num%5 != 0 {
			continue
		}
		fmt.Printf("sum: %v + %v := %v\n", sum, num, sum+num)
		sum = sum + num
	}
	fmt.Printf("sum: %v\n", sum)
}
