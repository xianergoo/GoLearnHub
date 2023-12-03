package main

import (
	"fmt"
	"math/rand"
)

// 洗牌算法
func shuffle(arr []int) {
	var i, j int
	// var temp int
	for i = len(arr) - 1; i > 0; i-- {
		j = rand.Intn(i)
		//fmt.Printf("j: %v\n", j)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		shuffle(intArr)
		fmt.Println(intArr)
	}
}
