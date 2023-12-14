package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler)

	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-id")

	// 检查是否存在会话ID
	if session.Values["authenticated"] != true {
		http.Redirect(w, r, "/login", 302)
		return
	}

	// 在此处处理受保护的页面
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// 在此处执行用户验证

	// 验证成功，创建session
	session, _ := store.Get(r, "session-id")
	session.Values["authenticated"] = true
	session.Save(r, w)

	// 重定向到首页
	http.Redirect(w, r, "/", 302)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-id")
	session.Values["authenticated"] = false
	session.Save(r, w)

	// 重定向到登录页面
	http.Redirect(w, r, "/login", 302)
}
