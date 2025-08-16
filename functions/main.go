package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// fmt.Println("Hello world")
	// result := addition(2, 4)

	// fmt.Println(result)
	// answer, _ := addder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// fmt.Println(answer)

	randomNumber()

}

func addition(val1 int, val2 int) int { // need to pass the type of variable
	return val1 + val2
}

func addder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "We can return two and multiple value like this"
}

func randomNumber() {
	// by math
	fmt.Println(rand.IntN(5))

	fmt.Println()

}
