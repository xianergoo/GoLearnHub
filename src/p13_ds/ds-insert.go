package main

import "fmt"

func insertSort(arr []int) {
	// 总循环次数，元素个数
	for i := 0; i < len(arr); i++ {
		// 从后面第二个开始，取出找到合适的位置（插入）
		for j := i; j > 0; j-- {
			// 比较
			if arr[j] > arr[j-1] {
				// 交换
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
	fmt.Println(arr)
}

func main() {
	values := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	insertSort(values)
}
