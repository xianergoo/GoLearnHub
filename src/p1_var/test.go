package main

import (
	"fmt"
)

func main() {
	var mystr01 string = "hello"
	var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("myStr01: %s\n", mystr01)
	fmt.Printf("myStr02: %s\n", mystr02)

	// var a int64
	// // var b uint
	// a = math.MaxInt64
	// fmt.Printf("%v\n", int64(a))

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)

	// d := [3]int{1, 2}
	// fmt.Println(a == d) //p1_var\test.go:24:19: invalid operation: a == d (mismatched types [2]int and [3]int)

	var array [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化数组中索引为 1 和 3 的元素
	// array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化数组中指定的元素
	// array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	for index, value := range array {
		fmt.Printf("index: %d, value: %v\n", index, value)
	}

	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{}
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)
	// fmt.Println(numListEmpty == numList) //invalid operation: numListEmpty == numList (slice can only be compared to nil)

	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6]
	//这打印出来长度为2
	fmt.Printf("myslice为 %v, 其长度为: %d, cap: %d\n", myslice, len(myslice), cap(myslice))
	myslice2 := myslice[:6]
	fmt.Printf("myslice: %v\n", myslice2)
	myslice = myslice[:cap(myslice)]
	fmt.Printf("myslice: %v\n", myslice)
	//为什么 myslice 的长度为2，却能访问到第四个元素
	fmt.Printf("myslice的第四个元素为: %d", myslice[5])

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	// copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println("%v\n", slice1)
	// fmt.Println("%v\n", slice1)
}
