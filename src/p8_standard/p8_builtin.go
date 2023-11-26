package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	i := append(s1, 4)
	fmt.Printf("i: %v\n", i)

	s2 := []int{3, 5, 6}
	i2 := append(s1, s2...)
	fmt.Printf("i2: %v\n", i2)

	s3 := "hello world"
	i3 := len(s3)
	fmt.Printf("i: %v\n", i3)

	s4 := []int{1, 2, 3}
	fmt.Printf("len(s2): %v\n", len(s4))

	name := "tom"
	print(name)

}
