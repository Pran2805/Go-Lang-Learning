package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	var url string = "https://jsonplaceholder.typicode.com/todos/1"

	res, err := http.Get(url)
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
