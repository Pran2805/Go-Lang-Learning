package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello")

	content := "This content need to go in file"

	file, err := os.Create("./file.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length is:", length)

	defer file.Close()
	readFile("./file.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)

	fmt.Println("Data", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
