package spiral_matrix_ii

import "fmt"

func generateMatrix(n int) [][]int {
	left, right := 0, n-1
	top, bottom := 0, n-1
	fornum := 1
	target_num := n * n
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for fornum <= target_num {
		for i := left; i <= right; i++ {
			res[top][i] = fornum
			fornum++
			fmt.Println(res)
		}
		top++
		for i := top; i <= bottom; i++ {
			res[i][right] = fornum
			fornum++
			fmt.Println(res)
		}
		right--
		for i := right; i >= left; i-- {
			res[bottom][i] = fornum
			fornum++
			fmt.Println(res)
		}
		bottom--
		for i := bottom; i >= top; i-- {
			res[i][left] = fornum
			fornum++
			fmt.Println(res)
		}
		left++
	}
	return res
}

func Add(a int, b int) int {
	return a + b
}

func generateMatrix2(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1
	tar := n * n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for num <= tar {
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}
	return matrix
}
