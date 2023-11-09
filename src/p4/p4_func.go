package main

import "fmt"

/*
func function_name ([parameter list]) [return_types]{
	function body
} */

func func1() {
	fmt.Println("no params no return")
}

func sum(a, b int) (ret int) {
	ret = a + b
	return ret

}

func fun_f2() (int, int) {
	return 1, 2
}

//try to modify func params

func func_parms(a int, b int) {
	a = 100
}

func func_args(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)

	}
}

func func_args_2(name string, male bool, args ...int) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("male: %v\n", male)
	for _, v := range args {
		fmt.Printf("v: %v\n", v)

	}
}

func main() {
	func1()
	f := sum(1, 2)
	fmt.Printf("f: %v\n", f)

	f2 := sum(fun_f2())
	fmt.Printf("f2: %v\n", f2)

	f3 := 10
	func_parms(f3, 1)
	fmt.Printf("f3: %v\n", f3)

	// func_parms(f3, 1)

	func_args(1, 2, 3, 4, 3)

	func_args_2("john", true, 1, 2, 3, 20)

	test_func("john", sayHello)

	ff := calc("+")
	r := ff(1, 2)
	fmt.Printf("r: %v\n", r)

	ff = calc("-")
	r = ff(2, 1)
	fmt.Printf("r: %v\n", r)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func calc(operator string) func(int, int) int {
	switch operator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func sayHello(name string) {
	fmt.Printf("hello: %v\n", name)
}

func test_func(name string, f func(string)) {
	f(name)
}
