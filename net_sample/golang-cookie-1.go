package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// 验证用户名和密码
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Println("username: %s, password: %s", username, password)
		if username == "admin" && password == "admin123" {
			// 设置记住我Cookie
			cookie := &http.Cookie{
				Name:    "remember_me",
				Value:   "1",
				Expires: time.Now().Add(30 * 24 * time.Hour),
			}
			http.SetCookie(w, cookie)

			fmt.Fprint(w, "Login successful.")
			return
		}

		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// 检查记住我Cookie
		cookie, err := r.Cookie("remember_me")
		if err == nil && cookie.Value == "1" {
			fmt.Fprint(w, "Welcome back!")
			return
		}

		// 显示登录页面
		fmt.Fprint(w, "Please login.")
	})

	http.ListenAndServe(":8080", nil)
}
