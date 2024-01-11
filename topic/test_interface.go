package main

import (
	"fmt"
)

func assert(i interface{}) {
	v, ok := i.(int)
	if ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println(i, ok)
	}

}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("我是字符串，我的值是：%s\n", i.(string))
	case int:
		fmt.Printf("我是int，我的值是：%d\n ", i.(int))
	default:
		fmt.Printf("未知类型\n")
	}
}

func main() {
	var s interface{} = 56
	assert(s)
	var i interface{} = "golang-teck-stack.com"
	assert(i)

	findType("golang-teck-stack.com")
	findType(100)
	findType(89.98)
}
