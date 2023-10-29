package main

import "fmt"

func getNameAndAge() (string, int) {
	return "tom", 20
}

func main() {

	//匿名变量
	//使用name, age := getNameAndAge()  age必须进行使用
	//使用 _ 可以不使用
	name, _ := getNameAndAge()
	fmt.Printf("name: %s\n", name)

	/* //短变量声明 :=
	//只能在函数内部使用短变量
	name := "tom"
	age := 20
	b := true
	fmt.Printf("name: %s,\n age: %d,\n b: %v\n", name, age, b)
	*/
	/* 	//批量初始化
	   	var name, age, b = "Tom", 20, true
	   	fmt.Printf("name: %s,\n age: %d,\n b: %v\n", name, age, b) */

	/* //类型推断
	var name = "sam"
	var age = 20
	var male = true

	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	if male {
		fmt.Printf("sex is male\n")
	} else {
		fmt.Printf("sex is female\n")
	} */

}
