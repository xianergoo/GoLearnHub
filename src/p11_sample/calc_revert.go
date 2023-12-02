package main

import "fmt"

func main() {
	var s string
	fmt.Printf("input a string\n")
	fmt.Scanf("%s\n", &s)
	reverts := revert(s)
	fmt.Printf("reverts: %v\n", reverts)

}

func revert(s string) (reverts string) {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	reverts = string(a)
	return
}
