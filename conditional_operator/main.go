package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func switch_case() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")
		fallthrough
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("Dice value is 3")
	case 4:
		fmt.Println("Dice value is 4")
	case 5:
		fmt.Println("Dice value is 5")
	case 6:
		fmt.Println("Dice value is 6")
	default:
		fmt.Println("Invalid value", diceNumber)
	}
}
func main() {
	// ifElse()
	switch_case()
}
