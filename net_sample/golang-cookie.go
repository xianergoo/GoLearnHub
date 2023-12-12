package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 获取请求中的Cookie
		cookie, err := r.Cookie("username")
		if err == nil {
			fmt.Fprintf(w, "Hello, %s", cookie.Value)
			return
		}

		// 设置响应的Cookie
		cookie = &http.Cookie{
			Name:  "username",
			Value: "Alice",
		}
		http.SetCookie(w, cookie)

		fmt.Fprint(w, "Welcome, new user!")
	})

	http.ListenAndServe(":8080", nil)
}
