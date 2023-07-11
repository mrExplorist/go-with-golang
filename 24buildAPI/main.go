package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Building API with Go

// Model for course - file


type Course struct {
	CourseId string `json:"course_id"`
	CourseName string `json:"course_name"`
	CoursePrice int `json:"course_price"` // - in strings will not be included in the JSON output 
	Author *Author `json:"author"`
}


// Model for user - file 

type Author struct {
	FullName string `json:"full_name"`
	Website string `json:"website"`
}



// fake DB

var courses []Course // slice of course


// middleware , helper - file

func (c * Course) IsEmpty() bool {

	return c.CourseId == "" && c.CourseName == "" && c.CoursePrice == 0 && c.Author == nil
	
	}



func main(){
	fmt.Println("Welcome to the course API , listening on port 8080")

	r := mux.NewRouter() // create a new router

	// seeding the fake DB with some courses data 

	courses = append(courses , Course{CourseId: "1" , CourseName: "Go" , CoursePrice: 100 , Author: &Author{FullName: "John Doe" , Website: "https://johndoe.com"}})
	courses = append(courses , Course{CourseId: "2" , CourseName: "Python" , CoursePrice: 200 , Author: &Author{FullName: "Jane Doe" , Website: "https://janedoe.com"}})
	courses = append(courses , Course{CourseId: "3" , CourseName: "Java" , CoursePrice: 300 , Author: &Author{FullName: "Jack Doe" , Website: "https://jackdoe.com"}})
	courses = append(courses , Course{CourseId: "4" , CourseName: "JavaScript" , CoursePrice: 400 , Author: &Author{FullName: "Jill Doe" , Website: "https://jilldoe.com"}})
	courses = append(courses , Course{CourseId: "5" , CourseName: "C#" , CoursePrice: 500 , Author: &Author{FullName: "James Doe" , Website: "https://jamesdoe.com"}})



	// routing
	r.HandleFunc("/" , serveHome).Methods("GET") // serve home route
	r.HandleFunc("/api/courses" , getAllCourses).Methods("GET") // get all courses
	r.HandleFunc("/api/course/{id}" , getSingleCourse).Methods("GET") // get a single course
	r.HandleFunc("/api/course" , createOneCourse).Methods("POST") // create a course
	r.HandleFunc("/api/course/{id}" , updateOneCourse).Methods("PUT") // update a course
	r.HandleFunc("/api/course/{id}" , deleteOneCourse).Methods("DELETE") // delete a course







	// listen to a port
	log.Fatal(http.ListenAndServe(":8080" , r))
	// syntax of ListenAndServe func 
	// func ListenAndServe(addr string, handler Handler) error
	// addr is the port number
	// handler is the router













}

// controllers - file

// serve home route

// syntax of handler func
// func (w http.ResponseWriter, r *http.Request) {} 


func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to the home page</h1>"))

}


func getAllCourses(w http.ResponseWriter, r *http.Request) {


	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json") // set header to json format for the response body and status code 200 by default 

	
	json.NewEncoder(w).Encode(courses) // encode courses to json and send it to the client
}


func getSingleCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get single course")
	w.Header().Set("Content-Type", "application/json")

	// get the id from the request query

	params := mux.Vars(r) // get the params from the request

	// params is a map of string keys and string values
	// loop through the courses and find the course with the id from the request query parameters 

	for _ , course := range courses { // loop through the courses slice 
		// if the course id matches the id from the request query parameters



		if course.CourseId == params["id"]{

			// send the course to the client

			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// if no course is found with the id from the request query parameters
	json.NewEncoder(w).Encode("No course found with the id provided")
	
}

// create a course
func createOneCourse(w http.ResponseWriter , r *http.Request){

	fmt.Println("Create a course")
	w.Header().Set("Content-type" , "application/json")
	// what if: the request body is empty
	if(r.Body == nil){
		json.NewEncoder(w).Encode("Please provide a course data to create")
		return
	}
	// what if: the request body is - { }
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) // decode the request body to a course struct
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please provide a valid course data to create")
		return

	}


	// check if requesting new courseName already exists in the courses slice and return an error if it does

	for _ , c := range courses {
		if c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course already exists")
			return
		}
	}
	




	// generate a random unique id for the course and convert it to string
	// add the course to the courses slice

	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	// send the course to the client
	json.NewEncoder(w).Encode(course)
	return
}

// update a course
func updateOneCourse(w http.ResponseWriter , r *http.Request){
	fmt.Println("Update a course")
	w.Header().Set("Content-type" , "application/json")
	// grab the id from the request query parameters
	params := mux.Vars(r)

	// loop , id , remove , add with my ID

	//... is known as the variadic parameter syntax and it allows you to pass a slice as individual elements to the append function.

	for index , course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index] , courses[index+1:]...) // append function in Go to remove an course from a courses slice at a specific index 
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses , course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: if no course is found with the id from the request query parameters
}

// delete a course
func deleteOneCourse(w http.ResponseWriter , r *http.Request){

	fmt.Println("Delete a course")
	w.Header().Set("Content-type" , "application/json")
	// grab the id from the request query parameters
	params := mux.Vars(r) 


	for index , course := range courses { // loop through the courses slice
		if course.CourseId == params["id"] {
			courses = append(courses[:index] , courses[index+1:]...) // append function in Go to remove an course from a courses slice at a specific index 
			
			json.NewEncoder(w).Encode("The course is deleted successfully")
			break
	}
}

}



