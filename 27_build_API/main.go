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

// Model for course - file

type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fale db
var courses []Course

// middleware or helper -file
func IsEmpty(c *Course) bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Build API in Golang")
	r:=mux.NewRouter()
	r.HandleFunc("/",serverRoot).Methods("GET")
	r.HandleFunc("/allcourses",getAllCourses).Methods("GET")
	r.HandleFunc("/onecourse/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/createcourse",createOneCourse).Methods("POST")
	r.HandleFunc("/update/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/delete/{id}",deleteOneCourse).Methods("DELETE")

	courses=append(courses,Course{CourseID:"2",CourseName: "Reactjs",CoursePrice: 299,Author: &Author{Fullname: "Kamalesh",Website:"kamy.tech"} })
	courses=append(courses,Course{CourseID:"5",CourseName: "Mern",CoursePrice: 19999,Author: &Author{Fullname: "Kamalesh",Website:"kamy.tech"} })
log.Fatal(http.ListenAndServe(":4000",r))
}

//controllers -file
// server home route

func serverRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Kamalesh</h1"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through the courses find matching id and return thr response

	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course ID found.")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	//what if :body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please provide a body with data for this operation!")
	}

	//what if :body is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	// if courses.isEmpty{
	// 	json.NewEncoder(w).Encode("{\"error\":\"please provide valid input!\"}")
	// 	return
	// }

	//generate a unique ID and convert to string
	//append the new course in courses

	rand.Seed(time.Now().UTC().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	//first grab id from req
	params := mux.Vars(r)

	//loop through the val remove the index and add the operation my ID
	for index,course:=range courses{
		if course.CourseID==params["id"]{
			courses=append(courses[:index],courses[index+1:]...)
			var course Course
			_=json.NewDecoder(r.Body).Decode(&course)
			course.CourseID=params["id"]
			courses=append(courses,course)
			json.NewEncoder(w).Encode(course)
			return
		}

}
//TODo when id not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index,course:=range courses{
	if course.CourseID == params["id"]{
		courses=append(courses[:index],courses[index+1:]...)
		break
	
}
	}
}