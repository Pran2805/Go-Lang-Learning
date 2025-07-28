package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const myUrl string = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	handling_web_request()
	handling_url()
}

func handling_web_request() {
	res, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res.Status))
	fmt.Println(res.StatusCode)
	fmt.Println(string(res.Proto))
	// fmt.Println(res)

	defer res.Body.Close()

	// read the response
	dataBytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataBytes))
}

func handling_url() {
	//parsing url
	res, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Println(res.Scheme) // mongodb / https /
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.RawQuery)

}
