package main

import "fmt"

func selectSort(nums []int) {

	for i := 0; i < len(nums)-1; i++ {
		min_index := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min_index] {
				min_index = j
			}
		}

		nums[i], nums[min_index] = nums[min_index], nums[i]
	}
	fmt.Printf("nums: %v\n", nums)
}

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	selectSort(values)
}
