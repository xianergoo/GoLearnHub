package main

import (
	"bufio"
	"flag"
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

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)

	}
}
func main() {
	// test1()
	// test2()
	fmt.Println(os.Getwd())

	// test3()
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "read frin %s faild, err: %v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}

}
