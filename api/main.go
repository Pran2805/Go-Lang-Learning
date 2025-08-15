package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses - should save in seperate file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake database
var courses []Course

// middleware - seperate file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}
func main() {
	fmt.Println("Creating API's")
	r := mux.NewRouter()

	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "ReactJs",
		CoursePrice: 499,
		Author: &Author{
			Fullname: "Pranav",
			Website:  "website",
		},
	})
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "ExpressJs",
		CoursePrice: 399,
		Author: &Author{
			Fullname: "Pranav",
			Website:  "another website",
		},
	})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourseById).Methods("GET")

	// another
	r.HandleFunc("/course/add", addCourse).Methods("POST")
	r.HandleFunc("/course/update/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/delete/{id}", updateCourse).Methods("DELETE")

	// listen
	log.Fatal(http.ListenAndServe(":8000", r))
}

// controllers - seperate file and folder
// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to HomePage</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Course by Id")
	w.Header().Set("Content-Type", "application/json")
	// get id
	params := mux.Vars(r)

	// find id in array and return the res
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id")
	return

}

func addCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Course")
	w.Header().Set("Content-Type", "application/json")

	// req.body -> empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Fields are empty")
		return
	}

	// req.body ->{}
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Fields are empty")
		return
	}

	// generate unique id and convert into string
	// append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course")
	w.Header().Set("Content-Type", "application/json")

	// first - get id from request
	params := mux.Vars(r)

	// var courses Course

	// loop, find course via id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Id is not Found")
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
