package main

import "fmt"

func twoSum(nums []int, target int) []int {
	targets := make(map[int]int)
	for i, value := range nums {
		targetIndx, ok := targets[target-value]
		if ok {
			return []int{targetIndx, i}
		}
		targets[value] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18

	fmt.Printf("Result: %v\n", twoSum(nums, target))
}
