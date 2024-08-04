package main

import (
	"fmt"
	"math"
)

// 暴力法
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	res := math.MaxInt32
	sublength := 0
	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				sublength = j - i + 1
				res = min(sublength, res)
				break
			}
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

// 滑动窗口
func minSubArrayLen2(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	res := math.MaxInt32
	sublength := 0
	i := 0

	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		fmt.Println("sum", sum)
		for sum >= target {

			sublength = (j - i + 1)
			res = min(res, sublength)
			fmt.Println(res)

			sum -= nums[i]
			i++
			fmt.Println("cur sum", sum)
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(nums)
	a := minSubArrayLen2(7, nums)
	fmt.Println(a)
}
