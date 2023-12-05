package main

import (
	"fmt"
	"sort"
)

// 二分查找循环实现
func BinarySearch(s []int, k int) int {
	// 低位高位 长度
	lo, hi := 0, len(s)-1
	// 如果低位小于等于高位一直循环
	for lo <= hi {
		// 取中间位，向右移动一位等于除以2
		m := (lo + hi) >> 1
		if s[m] < k {
			lo = m + 1
		} else if s[m] > k {
			hi = m - 1
		} else {
			return m
		}
	}
	return -1
}

// 二分查找递归调用
func BinarySearch2(arr *[]int, leftIndex int, rightIndex int, findVal int) {
	// 退出条件，判断条件不能为等号，因为在相等时仍然要进行一次判断。
	// 退出条件为等号可能会造成明明数组中有该值，却错过判断的情况发生。
	if leftIndex > rightIndex {
		fmt.Println("没找到")
		return
	}

	middleIndex := (leftIndex + rightIndex) / 2
	if findVal > (*arr)[middleIndex] {
		// 倘若不对middleIndex进行+1/-1的操作，函数有可能陷入死循环
		BinarySearch2(arr, middleIndex+1, rightIndex, findVal)
	} else if findVal < (*arr)[middleIndex] {
		BinarySearch2(arr, leftIndex, middleIndex-1, findVal)
	} else {
		fmt.Println("找到了！下标为：", middleIndex)
	}
}

func main() {
	// 二分查找必须有序
	arr := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	// 先排序
	sort.Ints(arr)
	fmt.Printf("arr: %v\n", arr)
	// 查找
	fmt.Println(BinarySearch(arr, 80))

	fmt.Println("---------------")

	BinarySearch2(&arr, 0, len(arr), 80)

}
