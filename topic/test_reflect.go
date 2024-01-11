// // package main

// // import (
// // 	"fmt"
// // 	"reflect"
// // )

// // func main() {
// // 	// var x int = 1
// // 	// fmt.Println("type: ", reflect.TypeOf(x))

// // 	// fmt.Println("Kind:  ", reflect.ValueOf(x).Kind())
// // 	// fmt.Println("Kind is Int? ", reflect.ValueOf(x).Kind() == reflect.Int)

// // 	var x float64 = 3.4
// // 	p := reflect.ValueOf(&x) // 获取x的地址
// // 	fmt.Println("settability of p: ", p.CanSet())
// // 	v := p.Elem()
// // 	fmt.Println("settability of v: ", v.CanSet())
// // 	v.SetFloat(7.1)
// // 	fmt.Println(v.Interface())
// // 	fmt.Println(x)

// // }

// // ----------------------------------------------
// // package main

// // import "fmt"

// // func say(text string) {
// // 	fmt.Println(text)
// // }

// // func main() {
// // 	var funcMap = make(map[string]interface{})
// // 	funcMap["say"] = say
// // 	funcMap["say"]("hello")
// // }

// //cannot call non-function funcMap["say"] (type interface {})
// // ----------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func say(text string) {
// 	fmt.Println(text)
// }

// func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value) {
// 	f := reflect.ValueOf(m[name])
// 	in := make([]reflect.Value, len(params))
// 	for k, param := range params {
// 		in[k] = reflect.ValueOf(param)
// 	}
// 	result = f.Call(in)
// 	return
// }

// func main() {
// 	var funcMap = make(map[string]interface{})
// 	funcMap["say"] = say
// 	Call(funcMap, "say", "hello")
// }

package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("Cat Eat")
}

func main() {
	var skills = []int{1, 2, 3, 4}
	fmt.Println(findPostion(skills, 3))
	fmt.Println(findPostion2(skills, 3))

	cat := Cat{}
	v := reflect.ValueOf(&cat)
}

func findPostion(arr []int, target int) int {
	for i, value := range arr {
		if value == target {
			return i
		}
	}
	return -1
}

func findPostion2(arr any, target any) int {
	fmt.Println(arr)
	arrs := reflect.ValueOf(arr)

	fmt.Println(arrs)

	for i := 0; i < arrs.Len(); i++ {
		v := arrs.Index(i)
		if v.Interface() == target {
			return i
		}
	}
	return -1
}
