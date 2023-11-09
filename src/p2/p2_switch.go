/* switch var1 {
case val1:
	...
case val2:
	...
default:
	...
}
*/

package main

import "fmt"

func main() {
	f()
	isweek()
	func_fall()
}

func func_fall() {
	a := 100
	switch a {
	case 100:
		fmt.Println("100")
		fallthrough
	case 200:
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	default:
		fmt.Println("other")
	}
}

func isweek() {
	day := 8
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("work day")
	case 6, 7:
		fmt.Println("week day")
	default:
		fmt.Println("invalid week days")
	}
}

func f() {
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("A")
	case "B":
		fmt.Println("B")
	default:
		fmt.Println("C")
	}
}
