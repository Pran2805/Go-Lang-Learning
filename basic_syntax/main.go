package main

import (
	"fmt"
)

// if the first letter is captial then it is the public otherwise private
const LoginToken string = "This is my variable token"

func main() {
	var username string = "Pranav"

	fmt.Println(username)
	fmt.Printf("Type of variable: %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("Type of variable: %T \n", isLoggedin)

	var smallVal uint8 = 255 // 255 is max value for uint8
	fmt.Println(smallVal)
	fmt.Printf("Type of variable: %T \n", smallVal)

	// aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Type of Variable: %T \n", anotherVariable)

	// implicit type
	var website = "hello world"
	fmt.Println(website)
	fmt.Printf("Type of Variable: %T \n", website)

	// no var style
	numberOfUser := 3000.12 // you can't use this style outside the method or function
	fmt.Println(numberOfUser)
	fmt.Printf("Type of Variable: %T \n", numberOfUser)

	fmt.Print(LoginToken)
}
