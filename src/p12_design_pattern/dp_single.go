package main

import (
	"fmt"
	"sync"
)

type Singleton interface {
	dosomething()
}

type singleton struct{}

func (this *singleton) dosomething() {
	fmt.Println("do some thing")
}

var instance1 *singleton

func GetInstance1() *singleton {
	if instance1 == nil {
		instance1 = &singleton{}
	}
	return instance1
}

var (
	once     sync.Once
	instance *singleton
)

func CreateSingleton() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	p1 := GetInstance1()
	fmt.Printf("p1: %p\n", p1)
	p2 := GetInstance1()
	fmt.Printf("p2: %p\n", p2)

	s1 := CreateSingleton()
	fmt.Printf("s1: %p\n", s1)

	s2 := CreateSingleton()
	fmt.Printf("s2: %p\n", s2)

}
