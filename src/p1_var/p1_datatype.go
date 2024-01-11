package main

import "fmt"

func func_test() {

}
func main() {
	var a rune
	/* var name string = "tom"
	age := 20
	b := true
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", b) */

	/* a := 100
	p := &a
	fmt.Printf("data: %v, type: %T", a, a)  //*int int指针类型
	*/

	/* a := [3]int{1, 2, 3}
	fmt.Printf("data: %v, type: %T", a, a) //int array
	*/

	/* a := []int{1, 2, 3}
	fmt.Printf("data: %v, type: %T", a, a) */ //动态数组

	fmt.Printf("data: %v, type: %T", func_test, func_test) // data: 0xd45200, type: func()

}
