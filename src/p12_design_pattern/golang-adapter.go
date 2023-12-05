package main

import "fmt"

type OldInterface interface {
	OldMethod()
}

type OldImpl struct {
}

func (OldImpl) OldMethod() {
	fmt.Println("旧方法实现")
}

type NewInterface interface {
	NewMethod()
}

type Adapter struct {
	OldInterface
}

// 适配器
func (a *Adapter) NewMethod() {
	fmt.Println("新方法实现")
	a.OldMethod()
}

func main() {
	oldInteface := OldImpl{}
	a := Adapter{OldInterface: oldInteface}
	a.NewMethod()
}
