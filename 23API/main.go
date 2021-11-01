package main

import (                               
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// course model -file
type Course struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  int     `json:"price"`
	Author *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fakedb
var courses []Course

// middlewares, helper - file
func (c *Course) IsEmpty() bool {
	return (c.Name == "")
}

// controllers -file
// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses route hit")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get courses by id route hit")
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, course := range courses {
		if course.Id == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get courses by id route hit")
	w.Header().Set("Content-type", "application/json")
	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data!")
		// r.Body.Close()
		return
	}
	// - body = {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data!")
		// r.Body.Close()
		return
	}

	for _, c := range courses {
		if c.Name == course.Name {
			json.NewEncoder(w).Encode("Course already exist!")
			return
		}
	}
	// generate a uid -> toString()
	// append course in courses
	rand.Seed(time.Now().UnixNano())
	course.Id = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get courses by id route hit")
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	// get id, remove, add with id
	for index, course := range courses {
		if course.Id == params["id"] {
			var myCourse Course
			myCourse.Id = params["id"]
			json.NewDecoder(r.Body).Decode(&myCourse)
			courses = append(courses[:index], courses[index+1:]...)
			courses = append(courses, myCourse)
			json.NewEncoder(w).Encode(myCourse)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found!")
}

func deleteCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete courses by id route hit")
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	// get id, remove
	for index, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found!")
}

func main() {
	courses = append(courses, Course{
		"2", "ReactJS", 299, &Author{"Divyansh", "dnx.dev"},
	})
	courses = append(courses, Course{
		"4", "AngularJS", 299, &Author{"Divyansh", "dnx.dev"},
	})
	fmt.Println("Welcome to API")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/api/course/all", getAllCourses).Methods("GET")
	r.HandleFunc("/api/course", createCourse).Methods("POST")
	r.HandleFunc("/api/course/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/api/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/api/course/{id}", deleteCourseById).Methods("DELETE")
	log.Fatal(http.ListenAndServe("localhost:4000", r))
}
