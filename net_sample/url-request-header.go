// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	request, err := http.NewRequest("GET", "https://example.com", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
// 	request.Header.Set("Authorization", "Bearer xxxxxxxxxxxxxxxxxxxxxxx")
// 	client := http.DefaultClient
// 	res, err := client.Do(request)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer res.Body.Close()

// 	fmt.Println(res.Status)

// }
package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func main() {
	username := "your_username"
	password := "your_password"
	url := "https://api.example.com"

	clent := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	auth := username + ":" + password
	basicAuth := "Basic" + base64.StdEncoding.EncodeToString([]byte(auth))
	request.Header.Add("Authorization", basicAuth)

	response, err := clent.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println(response.Status)

}
