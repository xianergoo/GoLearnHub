package main

import "fmt"

type triangle struct {
	size float32
}

func (t triangle) perimerter() float32 {
	return t.size * 3

}

type square struct {
	length float32
	hight  float32
}

func (s square) area() float32 {
	return s.length * s.hight
}
func main() {
	t := triangle{3}
	fmt.Printf("t.perimerter(): %v\n", t.perimerter())

	s := square{2, 3.5}
	fmt.Printf("s.area(): %v\n", s.area())

}
