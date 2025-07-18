package main

import "fmt"

func pointers() {
	var one *int

	fmt.Println(one)
	myNumber := 21
	var pointToMyNumber = &myNumber
	fmt.Println(pointToMyNumber)
	fmt.Println(&pointToMyNumber)
	fmt.Println(*pointToMyNumber)

	*pointToMyNumber = *pointToMyNumber * 2
	fmt.Println(myNumber)
}
func main() {
	fmt.Println("Hello world")

	pointers()
}
