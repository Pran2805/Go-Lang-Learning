package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wait sync.WaitGroup // mostly pointer
var mut sync.Mutex      // mostly pointer

func main() {
	websiteList := []string{
		"https://pranav-portfolio.pages.dev",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		wait.Add(1)
		go func(url string) {
			defer wait.Done()
			if res, err := http.Get(url); err != nil {
				printSafe("OOPS Endpoint: " + url)
			} else {
				printSafe(fmt.Sprintf("%d Status Code for %s", res.StatusCode, url))
			}
		}(web)
	}

	wait.Wait()
}

// short mutex-protected print
func printSafe(msg string) {
	mut.Lock()
	defer mut.Unlock()
	fmt.Println(msg)
}
