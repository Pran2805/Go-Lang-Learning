package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func variables() {
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

	fmt.Println(LoginToken)

}

func userInput() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something here:")

	// comma ok || error ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your input:", input)
}

func conversion() {
	fmt.Println("Welcome to Our Website")
	fmt.Println("Please rate our website between 1 to 10")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks For rating: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}
}

func timeManipulation() {
	fmt.Println("Time library")

	currentTime := time.Now()
	fmt.Println(currentTime)

	fmt.Println(currentTime.Format("01-02-2006"))
	fmt.Println(currentTime.Format("01-02-2006 Monday"))
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday")) // this is a syntax we can't change the 01-02-2006 otherwise it will give wrong values same with time and day

	createdDate := time.Date(2005, time.May, 28, 21, 21, 21, 21, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

// if the first letter is captial then it is the public otherwise private
const LoginToken string = "This is my variable token"

func main() {

	// variables()
	// userInput()
	// conversion()
	timeManipulation()
}
