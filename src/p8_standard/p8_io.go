package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// os.File 同时实现了 io.Reader 和 io.Writer
// strings.Reader 实现了 io.Reader
// bufio.Reader/Writer 分别实现了 io.Reader 和 io.Writer
// bytes.Buffer 同时实现了 io.Reader 和 io.Writer
// bytes.Reader 实现了 io.Reader
// compress/gzip.Reader/Writer 分别实现了 io.Reader 和 io.Writer
// crypto/cipher.StreamReader/StreamWriter 分别实现了 io.Reader 和 io.Writer
// crypto/tls.Conn 同时实现了 io.Reader 和 io.Writer
// encoding/csv.Reader/Writer 分别实现了 io.Reader 和 io.Writer

// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

func testCopy() {
	//file1 //file2
	r := strings.NewReader("Hello World")
	// if _, err := io.CopyBuffer()
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// r := strings.NewReader("hello world")
	// buf := make([]byte, 20)
	// r.Read(buf)
	// fmt.Printf("buf: %v\n", string(buf))
	// testCopy()

	r := strings.NewReader("1234 2121312 442")
	p := make([]byte, 3)
	for {
		n, err := r.Read(p)
		fmt.Println(n)
		if err != nil {
			if err == io.EOF {
				log.Printf("eof error : %d\n", n)
				break
			}
			log.Printf("error :%s\n", err)
			os.Exit(2)
		}
		log.Printf("-%d %s\n", n, string(p[:n]))
	}

}
