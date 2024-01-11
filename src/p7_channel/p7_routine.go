package main

import (
	"fmt"
	"net/http"
	"time"
)

func CheckAPI(api string) {
	_, err := http.Get(api)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Success: %s is up and running!\n", api)
}

func main() {
	// start := time.Now()

	// apis := []string{
	// 	"https://management.azure.com",
	// 	"https://dev.azure.com",
	// 	"https://api.github.com",
	// 	"https://outlook.office.com",
	// 	"https://graph.microsoft.com",
	// }

	// for _, api := range apis {
	// 	go CheckAPI(api)
	// }

	// time.Sleep(3 * time.Second)

	// elapsed := time.Since(start)
	// fmt.Printf("Done, It took %v seconds\n", elapsed.Seconds())
	// for i := 0; i < 7; i++ {
	// 	hello(i)
	// }
	go hello()
	fmt.Println("Goruntine End")
	time.Sleep(1 * time.Second)

}

func hello() {
	fmt.Println("Hello Golang Goruntine!")
}
