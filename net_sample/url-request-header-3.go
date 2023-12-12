package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		fmt.Printf("User-Agent: %s\n", userAgent)

		s := "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
		fmt.Println(s)

		// 根据不同的User-Agent字段做出相应的响应
		if userAgent == "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36" {

			fmt.Fprintf(w, "Hello, Chrome user!")
		} else if userAgent == "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36" {
			fmt.Fprintf(w, "Hello, Chrome user!")
		} else if userAgent == "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0" {
			fmt.Fprintf(w, "Hello, Firefox user!")
		} else {
			fmt.Fprintf(w, "Hello, unknown user!")
		}
	})

	http.ListenAndServe(":8080", nil)
}
