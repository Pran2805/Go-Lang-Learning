package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func array() {
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Cherry"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Mushroom"}

	fmt.Println("Vegetable List", vegList)
}

func slices() {
	var fruitList = []string{"Apple", "Banana"} // it is a splice
	fmt.Printf("Type of Fruitlist: %T \n", fruitList)
	fmt.Println("Fruitlist", fruitList)

	fruitList = append(fruitList, "Cherry")
	fmt.Println("Fruitlist", fruitList)

	// slicing of slice
	fruitList = append(fruitList[1:]) // exclusive for last element
	fmt.Println("Fruitlist", fruitList)

	highScore := make([]int, 4) // another way of declaration

	highScore[0] = 14
	highScore[1] = 2143
	highScore[2] = 5431
	highScore[3] = 413
	fmt.Println(highScore)
	highScore = append(highScore, 555, 666, 777)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "python", "golang", "reactnative"}
	fmt.Println(courses)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the index which you want to remove from the slice")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	index, _ := strconv.ParseInt(input, 10, 0)

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

func maps() {

}
func main() {
	// pointers()
	// array()
	// slices()
	maps()
}
