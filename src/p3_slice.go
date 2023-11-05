package main

import "fmt"

//var identifier []type

/* 	与数组一样，切片也是 Go 中的一种数据类型，它表示一系列类型相同的元素。
不过，与数组更重要的区别是切片的大小是动态的，不是固定的。
切片是数组或另一个切片之上的数据结构。 我们将源数组或切片称为基础数组。 通过切片，可访问整个基础数组，也可仅访问部分元素。

切片只有 3 个组件：

1. 指向基础数组中第一个可访问元素的指针。 此元素不一定是数组的第一个元素 array[0]。
2. 切片的长度。 切片中的元素数目。
3. 切片的容量。 切片开头与基础数组结束之间的元素数目。*/

func main() {
	/* 	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	   	fmt.Println(months)
	   	fmt.Println("length: ", len(months))
	   	fmt.Println("Capacity: ", cap(months)) */

	//months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	// quarter1 := months[0:3]
	// quarter2 := months[3:6]
	// quarter3 := months[6:9]
	// quarter4 := months[9:12]
	// fmt.Println(quarter1, len(quarter1), cap(quarter1))
	// fmt.Println(quarter2, len(quarter2), cap(quarter2))
	// fmt.Println(quarter3, len(quarter3), cap(quarter3))
	// fmt.Println(quarter4, len(quarter4), cap(quarter4))

	/* 	请注意，切片的长度不变，但容量不同。 我们来了解 quarter2 切片。 声明此切片时，
	   	你指出希望切片从位置编号 3 开始，最后一个元素位于位置编号 6。 切片长度为 3 个元素，但容量为 9，
	   	原因是基础数组有更多元素或位置可供使用，但对切片而言不可见。
	   	例如，如果你尝试打印类似 fmt.Println(quarter2[3]) 的内容，会出现以下错误：
	   	panic: runtime error: index out of range [3] with length 3。 */

	//  切片容量仅指出切片可扩展的程度。 因此，你可从 quarter2 创建扩展切片，如下例所示：
	// quarter2 := months[3:6]
	// quarter2Extend := quarter2[:4]
	// fmt.Println("quarter2: ", quarter2, len(quarter2), cap(quarter2))
	// fmt.Println("quarter2Extend: ", quarter2Extend, len(quarter2Extend), cap(quarter2Extend))
	/*
		请注意在声明 quarter2Extended 变量时，无需指定初始位置 ([:4])。 执行此操作时，
		Go 会假定你想要切片的第一个位置。 你可对最后一个位置 ([1:]) 执行相同的操作。
		Go 将假定你要引用所有元素，直到切片的最后位置 (len()-1)。 */

	//.....

	/* 	Go 提供了内置函数 append(slice, element)，便于你向切片添加元素。
	   	将要修改的切片和要追加的元素作为值发送给该函数。 然后，
	   	append 函数会返回一个新的切片，将其存储在变量中。 对于要更改的切片，变量可能相同。
	   	让我们看一下追加进程在代码中的显示方式：
	*/
	// var numbers []int
	// for i := 0; i < 10; i++ {
	// 	/*
	// 		numbers[i] = i
	// 		panic: runtime error: index out of range [0] with length 0
	// 		尝试直接赋值报错，也就是说切片无法自动扩容需要使用append方法
	// 	*/
	// 	numbers = append(numbers, i)
	// 	fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	// }
	// fmt.Println("numbers: ", numbers)

	/*
		  output:
			0       cap=1   [0]
			1       cap=2   [0 1]
			2       cap=4   [0 1 2]
			3       cap=4   [0 1 2 3]
			4       cap=8   [0 1 2 3 4]
			5       cap=8   [0 1 2 3 4 5]
			6       cap=8   [0 1 2 3 4 5 6]
			7       cap=8   [0 1 2 3 4 5 6 7]
			8       cap=16  [0 1 2 3 4 5 6 7 8]
			9       cap=16  [0 1 2 3 4 5 6 7 8 9]

		此输出很有意思。 特别是对于调用 cap() 函数所返回的内容。
		一切看起来都很正常，直到第 3 次迭代，此时容量变为 4，切片中只有 3 个元素。
		在第 5 次迭代中，容量又变为 8，第 9 次迭代时变为 16。
		你注意到容量输出中的模式了吗？ 当切片容量不足以容纳更多元素时，
		Go 的容量将翻倍。 它将新建一个具有新容量的基础数组。
		无需执行任何操作即可使容量增加。 Go 会自动扩充容量。 需要谨慎操作。
		有时，一个切片具有的容量可能比它需要的多得多，这样你将会浪费内存。
	*/

	//删除项
	// letters := []string{"A", "B", "C", "D", "E"}
	// remove := 2

	// if remove < len(letters) {
	// 	fmt.Println("Before: ", letters, "Remove", letters[remove])
	// 	letters = append(letters[:remove], letters[remove+1:]...)
	// 	fmt.Println("After: ", letters)
	// }

	/* 	此代码会从切片中删除元素。
	   	它用切片中的下一个元素替换要删除的元素，
	   	如果删除的是最后一个元素，则不替换。 */

	//创建切片的副本
	// slice2 := make([]string, 3)
	// copy(slice2, letters[1:4])
	// fmt.Println("slice2: ", slice2)

	//...

	// letters := []string{"A", "B", "C", "D", "E"}
	// fmt.Println("Before", letters)

	// slice1 := letters[0:2]
	// slice2 := letters[1:4]
	// fmt.Println("slice1: ", slice1)
	// fmt.Println("slice2: ", slice2)

	// fmt.Println("slice1[1]: ", slice1[1])
	// slice1[1] = "Z"
	// fmt.Println("slice1[1]: ", slice1[1])

	// fmt.Println("After", letters)
	// fmt.Println("Slice2", slice2)

	/* 	请注意对 slice1 所做的更改如何影响 letters 数组和 slice2。
	   	可在输出中看到字母 B 已替换为 Z，它会影响指向 letters 数组的每个切片。
		若要解决此问题，你需要创建一个切片副本，它会在后台生成新的基础数组。 可以使用以下代码：
	*/

	//...
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice1 := letters[0:2]

	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])

	slice1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)

	slice_array := [5]int{1, 2, 3, 4, 6}
	slice3 := slice_array[0:5]
	fmt.Printf("slice3: %v, %d\nslice_array: %v, %d\n", slice3, len(slice3), slice_array, len(slice_array))

	slice3[0] = 5
	slice3 = append(slice3, 7)
	fmt.Printf("slice3: %v, %d\nslice_array: %v, %d\n", slice3, len(slice3), slice_array, len(slice_array))
}
