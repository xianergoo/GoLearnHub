package main

import "fmt"

func main() {
	// f6()
	f7()
}

func f6() {
	//  Monday Tuesday Wednesday Thursday Friday Saturday Sunday
	var c string
	fmt.Println("输入一个字符：")
	fmt.Scan(&c)

	if c == "S" {
		fmt.Println("输入第二个字符：")
		fmt.Scan(&c)
		if c == "a" {
			fmt.Println("Saturday")
		} else if c == "u" {
			fmt.Println("Sunday")
		} else {
			fmt.Println("输入错误")
		}
	} else if c == "F" {
		fmt.Println("Friday")
	} else if c == "M" {
		fmt.Println("Monday")
	} else if c == "T" {
		fmt.Println("输入第二个字符：")
		fmt.Scan(&c)
		if c == "u" {
			fmt.Println("Tuesday")
		} else if c == "h" {
			fmt.Println("Thursday")
		} else {
			fmt.Println("输入错误")
		}
	} else if c == "W" {
		fmt.Println("Wednesday")
	} else {
		fmt.Println("输入错误")
	}
}

func f5() {
	if score := 80; score >= 60 && score <= 70 {
		fmt.Println("C")
	} else if score > 70 && score <= 90 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}

func f7() {
	a, b, c := 100, 102, 3
	if a > b {
		if a > c {
			fmt.Println("a")
		} else {
			fmt.Println("c")
		}
	} else if b > c {
		fmt.Println("b")

	} else {
		fmt.Println("c")
	}
}
