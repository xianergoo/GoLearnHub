package main

import "fmt"

func maxSubArraySumWithIndices(nums []int) (int, []int) {
	n := len(nums)
	if n == 0 {
		return 0, nil
	}

	maxSum := nums[0]
	maxSofar := nums[0]
	startIdxs := make([]int, n)
	startIdxs[0] = 0
	maxStartIndex := 0
	for i := 1; i < n; i++ {
		if nums[i] > maxSofar+nums[i] {
			maxSofar = nums[i]
			startIdxs[i] = i
		} else {
			maxSofar = maxSofar + nums[i]
			startIdxs[i] = startIdxs[i-1]
		}

		if maxSofar > maxSum {
			maxSum = maxSofar
			maxStartIndex = startIdxs[i]
		}
	}

	fmt.Println(maxSum, maxStartIndex)
	currentSum := 0
	subarry := make([]int, 0)
	for i := maxStartIndex; i < n; i++ {
		currentSum = currentSum + nums[i]
		subarry = append(subarry, nums[i])
		if currentSum == maxSum {
			break
		}
	}
	return maxSum, subarry
}

func main() {
	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	sum, arr2 := maxSubArraySumWithIndices(arr)
	fmt.Println("Maximum contiguous sum is", sum, arr2)
}
