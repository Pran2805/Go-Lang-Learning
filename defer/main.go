package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello")

	// defer uses a last in first out lifo
}
