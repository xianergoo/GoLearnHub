package main

import "fmt"

const MAX int = 3

func main() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("a's location: %x\n", &a)
	fmt.Printf("ip's location: %x\n", ip)
	fmt.Printf("ip value: %v\n", *ip)

	ar := []int{1, 2, 4}
	var i int
	var ptr [MAX]*int
	fmt.Printf("ptr: %v\n", ptr)
	for i = 0; i < MAX; i++ {
		ptr[i] = &ar[i]
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("ar[%d] = %d\n", i, *ptr[i])
	}

}
