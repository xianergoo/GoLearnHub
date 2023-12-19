// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"strings"
// )

// func jsonFormat(jsonbyte []byte) string {
// 	var jsonString bytes.Buffer
// 	if err := json.Indent(&jsonString, []byte(jsonbyte), "", " "); err != nil {
// 		return ""
// 	}
// 	return jsonString.String()
// }

// func testGet() {
// 	// https://www.juhe.cn/box/index/id/73
// 	url := "http://apis.juhe.cn/simpleWeather/query?key=087d7d10f700d20e27bb753cd806e40b&city=北京"
// 	r, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer r.Body.Close()
// 	b, _ := io.ReadAll(r.Body)
// 	fmt.Printf("b: %v\n", jsonFormat(b))
// }

// func testGet2() {
// 	params := url.Values{}
// 	url, err := url.Parse("http://apis.juhe.cn/simpleWeather/query")
// 	if err != nil {
// 		return
// 	}
// 	params.Set("key", "087d7d10f700d20e27bb753cd806e40b")
// 	params.Set("city", "shanghai")
// 	url.RawQuery = params.Encode()
// 	urlPath := url.String()
// 	fmt.Println(urlPath)

// 	resp, err := http.Get(urlPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer resp.Body.Close()

// 	body, _ := io.ReadAll(resp.Body)
// 	fmt.Println(jsonFormat(body))
// }

// func testParseJson() {
// 	type result struct {
// 		Args    string            `json:"args"`
// 		Headers map[string]string `json:"headers"`
// 		Origin  string            `json:"origin"`
// 		Url     string            `json:"url"`
// 	}

// 	resp, err := http.Get("http://httpbin.org/get")
// 	if err != nil {
// 		return
// 	}
// 	defer resp.Body.Close()
// 	body, _ := io.ReadAll(resp.Body)
// 	fmt.Println(string(body))
// 	var res result
// 	_ = json.Unmarshal(body, &res)
// 	fmt.Printf("%#v", res)
// }

// func testAddHeader() {
// 	client := &http.Client{}
// 	req, _ := http.NewRequest("GET", "http://httpbin.org/get", nil)
// 	req.Header.Add("name", "老郭")
// 	req.Header.Add("age", "80")
// 	resp, _ := client.Do(req)
// 	body, _ := io.ReadAll(resp.Body)
// 	fmt.Printf(string(body))
// }

// func testPost() {
// 	path := "http://apis.juhe.cn/simpleWeather/query"
// 	urlValues := url.Values{}
// 	urlValues.Add("key", "087d7d10f700d20e27bb753cd806e40b")
// 	urlValues.Add("city", "上海")
// 	r, err := http.PostForm(path, urlValues)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer r.Body.Close()
// 	b, _ := io.ReadAll(r.Body)
// 	fmt.Printf("b: %v\n", jsonFormat(b))
// }

// func testPost2() {
// 	urlValues := url.Values{
// 		"name": {"老郭"},
// 		"age":  {"80"},
// 	}
// 	reqBody := urlValues.Encode()
// 	resp, _ := http.Post("http://httpbin.org/post", "text/html", strings.NewReader(reqBody))
// 	body, _ := io.ReadAll(resp.Body)
// 	fmt.Println(jsonFormat(body))
// }

// func testPostJson() {
// 	data := make(map[string]string)
// 	data["site"] = "www.duoke360.com"
// 	data["name"] = "多课网"
// 	bytesData, _ := json.Marshal(data)
// 	resp, _ := http.Post("http://httpbin.org/post", "application/json", bytes.NewReader(bytesData))
// 	body, _ := io.ReadAll(resp.Body)
// 	fmt.Println(string(body))
// }

// func main() {
// 	// testGet()
// 	// testGet2()
// 	// testParseJson()
// 	// testAddHeader()
// 	// testPost()
// 	// testPost2()
// 	testPostJson()
// }
