package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func testReadDir() {
	fi, _ := ioutil.ReadDir(".")
	for _, v := range fi {
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}

func testWriteFile() {
	ioutil.WriteFile("a.txt", []byte("hello world"), 0664)
}

func testCreateTempFile() {
	content := []byte("temporary file's content")
	tmpFile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tmpFile: %v\n", tmpFile.Name())

	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// r := strings.NewReader("hello world")
	// b, err := ioutil.ReadAll(r)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("b: %v\n", string(b))

	// f, _ := os.Open("a.txt")
	// defer f.Close()

	// b, err := ioutil.ReadAll(f)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("b: %v\n", string(b))

	// testReadDir()
	// testWriteFile()
	testCreateTempFile()
}
