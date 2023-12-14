package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-id")

	// 检查session中是否有userID键
	if session.Values["userID"] == nil {
		// 如果没有，说明用户是新的，将userID存储在session中
		session.Values["userID"] = "user-" + fmt.Sprint(time.Now().UnixNano())
		session.Save(r, w)
	}

	// 获取当前在线用户数
	onlineCount := getOnlineCount()

	// 显示在线用户数
	fmt.Fprintf(w, "Online users: %d", onlineCount)
}

func getOnlineCount() int {
	// 遍历所有session，统计在线用户数
	count := 0

	store := sessions.NewFilesystemStore("", []byte("secret-key"))

	store.Walk(func(sess *sessions.Session, err error) error {
		if err != nil {
			return err
		}
		if sess.Values["userID"] != nil {
			count++
		}
		return nil
	})

	return count
}
