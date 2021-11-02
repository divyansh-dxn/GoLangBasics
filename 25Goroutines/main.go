package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	websites := []string{
		"https://google.com",
		"https://github.com",
		"https://fb.com",
	}
	for _, web := range websites {
		wg.Add(1)
		getStatus(web)
	}
	wg.Wait()
}

func getStatus(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.StatusCode, "for endpoint", endpoint)
	defer wg.Done()
}
