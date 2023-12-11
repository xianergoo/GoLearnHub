// package main

// import (
// 	"fmt"
// 	"net/url"
// )

// func main() {
// 	// 解析URL
// 	request_url := "https://www.example.com/search?q=golang#top"
// 	u, err := url.Parse(request_url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Printf("u.Scheme: %v\n", u.Scheme)
// 	fmt.Printf("u.Host: %v\n", u.Host)
// 	fmt.Printf("u.Path: %v\n", u.Path)
// 	fmt.Printf("u.Query(): %v\n", u.Query())
// 	fmt.Printf("u.Fragment: %v\n", u.Fragment)
// 	fmt.Printf("u.User: %v\n", u.User)
// 	fmt.Printf("u.RequestURI(): %v\n", u.RequestURI())

// }

package main

import (
	"fmt"
	"net/url"
)

func main() {
	// 构建URL
	u := &url.URL{
		Scheme: "https",
		Host:   "www.example.com",
		Path:   "/search",
	}
	q := u.Query()
	q.Set("q", "golang")
	q.Set("index", "5")
	u.RawQuery = q.Encode()

	// 输出URL字符串
	fmt.Println(u.String())
}
