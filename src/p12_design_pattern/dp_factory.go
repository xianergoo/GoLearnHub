package main

import "fmt"

type Fruit interface {
	grant()
	pick()
}
type Apple struct {
}

func (this *Apple) grant() {
	fmt.Println("grant apple")
}

func (this *Apple) pick() {
	fmt.Println("pick apple")
}

type Orange struct {
}

func (this *Orange) grant() {
	fmt.Println("grant Orange")
}

func (this *Orange) pick() {
	fmt.Println("pick Orange")
}

func NewFruit(t int) Fruit {
	switch t {
	case 1:

		return &Apple{}
	case 2:
		return &Orange{}
	}
	return nil
}

func main() {
	apple := NewFruit(1)
	apple.grant()
	apple.pick()
	fmt.Println("-----------------")
	orange := NewFruit(2)
	orange.grant()
	orange.pick()
}
