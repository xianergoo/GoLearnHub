package main

import (
	"net/http"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

// func main() {
// 	ghz := Person{"ghz", 80}
// 	muban := "hello, {{.Name}}, Your age {{.Age}}"
// 	tmpl, err := template.New("test").Parse(muban)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = tmpl.Execute(os.Stdout, ghz)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func tmpl(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("test.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/tmpl", tmpl)
	server.ListenAndServe()
}
