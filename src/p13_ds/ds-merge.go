package main

import (
	"fmt"
)

func mergerSort(arr []int, a, b int) {
	// 长度小于等于1，不用排序
	if b-a <= 1 {
		return
	}
	// 取中值
	c := (a + b) / 2
	// 递归调用
	// 排左边
	mergerSort(arr, a, c)
	// 排右边
	mergerSort(arr, c, b)
	// 左边切片
	arrLeft := make([]int, c-a)
	// 右边切片
	arrRight := make([]int, b-c)
	// 拷贝左边数据到左边切片
	copy(arrLeft, arr[a:c])
	// 拷贝右边数据到右边切片
	copy(arrRight, arr[c:b])
	i := 0
	j := 0
	for k := a; k < b; k++ {
		if i >= c-a {
			arr[k] = arrRight[j]
			j++
		} else if j >= b-c {
			arr[k] = arrLeft[i]
			i++
		} else if arrLeft[i] < arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}
}

func main() {
	arr := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Println(arr)
	mergerSort(arr, 0, len(arr))
	fmt.Println(arr)
}
