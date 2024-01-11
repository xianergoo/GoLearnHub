package main

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error
	// CheckData() bool
}

type file struct {
}

func (d *file) WriteData(data interface{}) error {
	fmt.Println("WriteData: ", data)
	return nil
}

// 这里使用指针类型和非指针类型是不同的
// 指针类型代表实现的是d实例的接口
// 非指针类型
// func (d file) WriteData(data interface{}) error {
// 	fmt.Println("WriteData: ", data)
// 	return nil
// }

func main() {
	// f := new(file)
	var writer DataWriter
	f := &file{}
	fmt.Printf("%T\n", f)
	writer = f

	writer.WriteData("data")

}
