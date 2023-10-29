package main

import (
	"fmt"
	"strings"
)

func main() {

	//str func
	s := "hello world"
	fmt.Printf("len: %v\n", len(s))
	fmt.Printf("strings.Split(s, \" \")[0]: %v\n", strings.Split(s, " ")[0])
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))
	fmt.Printf("strings.HasPrefix(s, \"hello\"): %v\n", strings.HasPrefix(s, "hello"))
	fmt.Printf("strings.HasSuffix(s, \"world！\"): %v\n", strings.HasSuffix(s, "world！"))
	fmt.Printf("strings.Index(s, \"l\"): %v\n", strings.Index(s, "l"))
	fmt.Printf("strings.LastIndex(s, \"l\"): %v\n", strings.LastIndex(s, "l"))

	//切片操作
	/* s := "hello world"
	a := 1
	b := 5
	fmt.Printf("s[a]: %s\n", s[a])
	fmt.Printf("s[a：b]: %s\n", s[a:b])
	fmt.Printf("s[a:]: %s\n", s[a:])
	fmt.Printf("s[:m]: %s\n", s[:b]) */

	//转义字符
	/* s1 := "hello"
	print(s1 + "\n,")
	print(s1 + "\t,")

	fmt.Println("")

	s2 := "c:\\program file\\a"
	print(s2) */

	//字符串连接
	// var buffer bytes.Buffer
	// buffer.WriteString("tom")
	// buffer.WriteString(",")
	// buffer.WriteString("20")

	// fmt.Printf("buffer.String(): %s\n", buffer.String())

	/* s1 := "tom"
	s2 := "20"
	// msg := s1 + s2
	msg1 := fmt.Sprintf("%s, %s", s1, s2) // msg1: tom, 20,
	msg2 := fmt.Sprint("%s, %s", s1, s2)  // msg2: %s, %stom20
	msg3 := strings.Join([]string{s1, s2}, ", ")
	fmt.Printf("msg1: %v ... msg2: %v ... msg3: %s", msg1, msg2, msg3) */

	/* var str1 string = "hello world"
	var html string = `
		<html>
			<head><title>hello golang</title>
		</html>
	`

	fmt.Printf("str1: %v\n", str1)
	fmt.Printf("html: %v\n", html) */
}
