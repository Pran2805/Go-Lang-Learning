package main

import "fmt"

func ifElse() {
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Admin"
	} else {
		result = "Hackeeeeer"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}
}

func main() {
	// ifElse()
}
