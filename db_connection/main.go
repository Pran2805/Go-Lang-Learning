package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pran2805/go-lang-learning/router"
)

// in root directory there will be only one go file, otherwise it will not run
func main() {
	fmt.Println("Server is Getting Started!")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":8000", r))

}
