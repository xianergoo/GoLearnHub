package main

import (
	"fmt"
	"net/http"
	"time"
)

func header_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=60")                                        // 设置缓存时间为60秒
	w.Header().Set("Expires", time.Now().Add(time.Minute).UTC().Format(http.TimeFormat)) // 设置过期时间为1分钟后
	w.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))            // 设置最后修改时间为当前时间

	// if ifModifiedSince := r.Header.Get("If-Modified-Since"); if ifModifiedSince != "" {
	// 	lastModified, err := time.Parse(http.TimeFormat, ifModifiedSince)
	// 	if err == nil && lastModified.UTC().Add(time.Minute).After(time.Now().UTC()) {
	// 		w.WriteHeader(http.StatusNotModified)
	// 		return
	// 	}
	// }

	if ifModifiedSince := r.Header.Get("If-Modified-Since"); ifModifiedSince != "" {
		lastModified, err := time.Parse(http.TimeFormat, ifModifiedSince)
		if err == nil && lastModified.UTC().Add(time.Minute).After(time.Now().UTC()) {
			w.WriteHeader(http.StatusNotModified) // 返回304状态码表示内容未修改
			return
		}
	}

	fmt.Fprintf(w, "Hello, world!")

}

func main() {
	http.HandleFunc("/", header_handler)
	http.ListenAndServe(":8080", nil)

}
