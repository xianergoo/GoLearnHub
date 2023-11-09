package main

import "fmt"

//type NewType Type

func main() {
	type MyInt int
	var i MyInt
	i = 100

	fmt.Printf("i: %v i: %T\n", i, i)

	// 类型别名定义
	type MyInt2 = int
	// i 其实还是int类型
	var i2 MyInt2
	i2 = 100
	fmt.Printf("i2: %v i2: %T\n", i2, i2)
}

/* 类型定义相当于定义了一个全新的类型，与之前的类型不同；
但是类型别名并没有定义一个新的类型，而是使用一个别名来替换之前的类型
类型别名只会在代码中存在，在编译完成之后并不会存在该别名
因为类型别名和原来的类型是一致的，所以原来类型所拥有的方法，
类型别名中也可以调用，但是如果是重新定义的一个类型，
那么不可以调用之前的任何方法 */
