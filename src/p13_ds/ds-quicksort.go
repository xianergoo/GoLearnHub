package main

import "fmt"

func quickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	middle := arr[0]
	var left []int
	var right []int

	for i := 1; i < length; i++ {
		if middle < arr[i] {
			right = append(right, []int{arr[i]}...)
		} else {
			left = append(left, []int{arr[i]}...)
		}
	}

	middle_s := []int{middle}

	left = quickSort(left)
	right = quickSort(right)

	result := append(append(left, middle_s...), right...)
	return result

}

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	result := quickSort(values)
	fmt.Printf("result: %v\n", result)
}
