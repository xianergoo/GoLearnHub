package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func test1() {
	r := strings.NewReader("hello world!")
	r2 := bufio.NewReader(r)
	s, _ := r2.ReadString('\n')
	fmt.Printf("s: %v\n", s)
}

func test2() {
	s := strings.NewReader("ABCDEFGHIGKLMNOPQRSTUVWXYZ1234567890")
	br := bufio.NewReader(s)
	p := make([]byte, 10)

	for {
		n, err := br.Read(p)
		if err == io.EOF {
			break
		} else {
			fmt.Printf("p: %v\n", string(p[:n]))
		}
	}
}

//WRITER

func test3() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE, 0777)
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString("hello world 1111")
	w.Flush()

	f2, _ := os.Open("a.txt")
	defer f.Close()

	r := bufio.NewReader(f2)
	s, _ := r.ReadString('\n')
	fmt.Println(s)

}
func main() {
	// test1()
	// test2()

	test3()
}
