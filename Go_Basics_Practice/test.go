package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex
var storage []string

func main() {

	websites := []string{
		"https://go.dev",
		"https://fb.com",
		"https://google.com",
	}

	for _, val := range websites {
		wg.Add(1)
		go getStatusCode(val)
	}
	wg.Wait()
	fmt.Println(storage)
}

func getStatusCode(url string) {
	defer wg.
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("opps something went wrong", err)
	} else {
		mut.
		storage= append(storage,url)
		fmt.Printf("%d status code %s\n", res.StatusCode, url)
	}

}