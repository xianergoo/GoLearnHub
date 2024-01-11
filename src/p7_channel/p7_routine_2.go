package main

import (
	"fmt"
)

func main() {
	// go func() {
	// 	defer fmt.Println("A.defer")
	// 	func() {
	// 		defer fmt.Println("B.defer")
	// 		// 结束协程
	// 		runtime.Goexit()
	// 		defer fmt.Println("C.defer")
	// 		fmt.Println("B")
	// 	}()
	// 	fmt.Println("A")
	// }()
	// var i int = 0
	// for {
	// 	fmt.Println("haha", i)
	// 	i++
	// 	time.Sleep(time.Second * 2)
	// }

	// runtime.GOMAXPROCS(4)
	// go a()
	// go b()
	// c()
	// time.Sleep(time.Second)

}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func c() {
	for i := 1; i < 10; i++ {
		fmt.Println("C:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
