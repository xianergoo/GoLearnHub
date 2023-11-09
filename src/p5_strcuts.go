package main

import "fmt"

type Person struct {
	name string
}

func (per Person) showPerson() {
	fmt.Printf("per: %p\n", &per)
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

func (per *Person) showPerson2() {
	fmt.Printf("per: %p\n", per)
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

func main() {
	p1 := Person{name: "tom"}
	fmt.Printf("p1: %p\n", &p1)
	p1.showPerson()
	fmt.Printf("p1: %v\n", p1)
	fmt.Println("---------------")
	p2 := &Person{name: "tom"}
	fmt.Printf("p2: %p\n", p2)
	p2.showPerson2()
	fmt.Printf("p2: %v\n", p2)
}
