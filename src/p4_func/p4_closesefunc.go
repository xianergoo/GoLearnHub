package main

import (
	"fmt"
	"strings"
)

func add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// callback function
func performOperation(numbers []int, operation func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = operation((num))
	}
	return result
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func jpg_func_normal(name string) string {
	if !strings.HasSuffix(name, ".jpg") {
		return name + ".jpg"
	}
	return name
}

func txt_func_normal(name string) string {
	if !strings.HasSuffix(name, ".txt") {
		return name + ".txt"
	}
	return name
}

func main() {

	//normal func get suffix
	fmt.Printf("jpg_func_normal(\"test0\"): %v\n", jpg_func_normal("test0"))
	fmt.Printf("txt_func_normal(\"test0\"): %v\n", txt_func_normal("test0"))

	// func factory1
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")

	fmt.Printf("jpgFunc(\"test\"): %v\n", jpgFunc("test"))
	fmt.Printf("txtFunc(\"test\"): %v\n", txtFunc("test"))

	// func factory2
	numbers := []int{1, 2, 3, 4, 5}
	// Square each number
	squared := performOperation(numbers, func(x int) int {
		return x * x
	})

	// double each number
	doubled := performOperation(numbers, 
		func(x int) int {
		return x * 2}
	)

	fmt.Printf("squared: %v\n", squared)
	fmt.Printf("doubled: %v\n", doubled)

	//闭包 = 函数 + 引用环境
	var f = add()
	fmt.Printf("f(10): %v\n", f(10))
	fmt.Printf("f(10): %v\n", f(10))

	fmt.Println(" ---------- ")
	f1 := add()
	fmt.Printf("f1(20): %v\n", f1(20))
	fmt.Printf("f1(20): %v\n", f1(30))

}
