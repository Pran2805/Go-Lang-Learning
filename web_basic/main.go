package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myUrl string = "http://localhost:8000"

func main() {
	// handling_web_request()
	// handling_url()
	// get_request()
	post_request()
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

func get_request() {
	res, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	fmt.Println("Status Code", res.StatusCode)
	fmt.Println("Content", res.ContentLength)

	// content, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("message:", string(content))

	// another way
	var resString strings.Builder
	content, _ := io.ReadAll(res.Body)
	byteContent, _ := resString.Write(content)
	fmt.Println("Message content:", byteContent)
	fmt.Println("Message:", resString.String())
}

func post_request() {
	// fake json data or payload
	requestBody := strings.NewReader(`
	{
		"name": "Pranav Shinde",
		"role": "author"
	}`)
	res, _ := http.Post(myUrl+"/post", "application/json", requestBody)
	// fmt.Println("res", res)
	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data))
	defer res.Body.Close()
}
