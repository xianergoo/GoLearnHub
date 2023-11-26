package main

import (
	"fmt"
	"net/http"
	"time"
)

func CheckAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		// ch <- fmt.Sprintf("Error: %v\n", err)
		ch <- fmt.Sprintf("Error: %s us down!\n", api)
		return
	}

	ch <- fmt.Sprintf("Success: %s is up and running!\n", api)
}

func main() {

	// ch <- x // sends (or writes ) x through channel ch
	// x = <-ch // x receives (or reads) data sent to the channel ch
	// <-ch // receives data, but the result is discarded

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://graph.microsoft.com",
	}

	start := time.Now()
	ch := make(chan string)
	for _, api := range apis {
		go CheckAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	timecost := time.Since(start)
	fmt.Printf("Done! it took %v second!\n", timecost.Seconds())
}
