package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(num int) string {
	switch {
	case num%15 == 0:
		return "FizzBuzz"
	case num%3 == 0:
		return "Fizz"
	case num%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(num)
}

// 要求用户输入一个数字，如果该数字为负数，则进入紧急状态
func f1() {
	val := 0
	for {

		fmt.Print("enter number:")
		fmt.Scanf("%d", &val)

		switch {
		case val < 0:
			panic("You entered a negative number!")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
	}
}

func f2() {
	val := 0
	fmt.Println("Enter number: ")
	fmt.Scanf("%d", &val)
	fmt.Println("entered: ", val)
}

func main() {
	f2()
	f1()
	/*
		for num := 1; num <= 100; num++ {
			fmt.Println(fizzbuzz(num))
		} */
}
