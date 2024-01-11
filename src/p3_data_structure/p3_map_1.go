package main

import "fmt"

func main() {
	notes := make(map[string]int)
	fmt.Println(len(notes))

	notes["a"] = 1
	fmt.Println(len(notes))

	s := "a"
	updateString(s)
	fmt.Println(s)
}

func updateString(s string) {
	s = "ssssssssupdate"
}
