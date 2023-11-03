package main

import "fmt"

type Employee struct {
	ID        int
	FirstName string
	Lastname  string
	Address   string
}

func main() {
	emp := Employee{}
	fmt.Printf("emp: %v\n", emp)
}
