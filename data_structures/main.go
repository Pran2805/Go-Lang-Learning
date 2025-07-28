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
	languages := make(map[string]string)

	languages["js"] = "Javascript"
	languages["go"] = "GoLang"
	languages["py"] = "Python"
	languages["ts"] = "Typescript"

	fmt.Println(languages)
	fmt.Println(languages["js"])

	delete(languages, "ts")
	fmt.Println(languages)

	// loops
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func structures() {
	// init struct
	pranav := User{"Pranav", "pranavshinde.as@gmail.com", true, 20}
	fmt.Println(pranav)
	fmt.Printf("%+v", pranav)
	fmt.Println(pranav.Name)
	pranav.getEmail()
	pranav.getEmail()
}
func main() {
	// pointers()
	// array()
	// slices()
	// maps()
	structures() // no oop in go lang
}

func (u User) getEmail() {
	fmt.Println(u.Email)
}
