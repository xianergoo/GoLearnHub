package main

import "fmt"

/* type Pet interface {
	eat()
}

type Dog struct {
}
type Cat struct {
}

func (cat Cat) eat() {
	fmt.Println("cat eat...")
}

func (dog Dog) eat() {
	fmt.Println("dog eat...")
} */

type Flyer interface {
	fly()
}
type Swimmer interface {
	swim()
}

type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct {
}

func (fish Fish) fly() {
	fmt.Println("fly...")
}

func (fish Fish) swim() {
	fmt.Println("swim...")
}

//OCP 开闭原则 可拓展，不可修改

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
	name string
	age  int
}

func (dog Dog) eat() {
	fmt.Println("dog eat...")
}

func (dog Dog) sleep() {
	fmt.Println("dog sleep...")
}

type Cat struct {
	name string
	age  int
}

func (cat Cat) eat() {
	fmt.Println("cat eat...")
}

func (cat Cat) sleep() {
	fmt.Println("cat sleep...")
}

type Person struct {
	name string
}

func (per Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {

	//OCP
	dog := Dog{}
	cat := Cat{}
	per := Person{}

	per.care(dog)
	per.care(cat)

	fmt.Println(".....................")

	var p Pet
	p = Cat{}
	p.eat()
	p = Dog{}
	p.eat()

	fmt.Println(".....................")

	var ff FlyFish
	ff = Fish{}
	ff.fly()
	ff.swim()
}
