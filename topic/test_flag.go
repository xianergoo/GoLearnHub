package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Int("age", 20, "请输入一个年龄：")
	flag.String("name", "老郭", "请输入一个姓名")
	flag.Parse()

	fmt.Printf("flag.Args(): %v\n", flag.Args())
	// for i := 0; i < flag.Args(); i++ {
	// 	s := flag.Arg(i)
	// 	fmt.Printf("s: %v\n", s)
	// }

	for i := 0; i < flag.NArg(); i++ {
		s := flag.Arg(i)
		fmt.Printf("s: %v\n", s)
	}
}
