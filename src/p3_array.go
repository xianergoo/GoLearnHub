package main

import "fmt"

//var variable_name [SIZE] variable_type

func main() {

	// 声明数组
	// var a [3]int
	// a[1] = 10
	// fmt.Println(a[0])
	// fmt.Println(a[1])
	// fmt.Print(a[len(a)-1])

	// 初始化数组
	/* 	cities := [5]string{"New York", "Paris", "Berlin", "Beijing"}
	   	fmt.Println("Cites: ", cities) */

	//数组中的省略号
	//如果你不知道你将需要多少个位置，但知道数据元素集，那么还有一种声明和初始化数组的方法是使用省略号 (...)，如下例所示：
	//q := [...]int{1, 2, 3}
	cities := [...]string{"New York", "Paris", "Berlin", "Beijing", "Shanghai"}
	fmt.Println("Cites: ", cities)

	//你能看出区别吗？ 末尾没有空字符串。 数组长度由你初始化它时输入的字符串决定。 如果你不再需要，则不保留你不知道的内存。
	//另一种有趣的数组初始化方法是使用省略号并仅为最后一个位置指定值。 例如，使用以下代码：

	numbers := [...]int{99: -1}
	fmt.Println("First: ", numbers[0])
	fmt.Println("Last: ", numbers[99])
	fmt.Println("Length: ", len(numbers))
	// a := -1
	// fmt.Println("-1: ", numbers[a])

	//多维数组
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("row: ", i, twoD[i])
	}
	fmt.Println("All data: ", twoD)

}
