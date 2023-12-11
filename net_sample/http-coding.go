// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	resp, err := http.Get("http://www.example.com/")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(body))
// }

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://www.example.com/login"
	data := []byte(`{"username": "admin", "password": "password"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
