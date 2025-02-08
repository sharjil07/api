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

type Course struct {
	CourseName  string  `json:"coursename"`
	CourseId    string  `json:"courseid"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:fullname`
	Website  string `json:"website"`
}

var courses []Course


func (c *Course) IsEmpty() bool {
	return c.CourseName == "" 
}

func main() {
    fmt.Println("Welcome to the API series")

	r:=mux.NewRouter()

	courses=append(courses, Course{CourseId: "2",CourseName: "ReactJs",
	CoursePrice: 299,Author: &Author{FullName: "Md Sharjil Alam",Website: "www.lerancodeonline.com"},})

	courses=append(courses, Course{CourseId: "3",CourseName: "NextJs",
	CoursePrice: 199,Author: &Author{FullName: "Danish Khan",Website: "www.lerancode.com"},})

	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllcourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOnecourse).Methods("GET")
	r.HandleFunc("/course",CreateoneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteAllcourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000",r))

}

func serveHome(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte ("<h1>Welcome to the API </h1>"))
}

func getAllcourses(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Get All The Courses ")
	w.Header().Set("Content-Type:","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOnecourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Get the Course ")
	w.Header().Set("Content-Type:","application/json")

	params:=mux.Vars(r)

	for _,course:=range courses{
		if course.CourseId==params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Can't find the course with that id ")
    return
	
}

func CreateoneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Create the Course ")
	w.Header().Set("Content-Type:","application/json")

	if r.Body==nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	var course Course

	_=json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}
	for _,c:=range courses{
		if c.CourseName== course.CourseName{
			json.NewEncoder(w).Encode("Already exist") 
			return
		}
		
	}

	rand.Seed(time.Now().UnixNano())

	course.CourseId=strconv.Itoa(rand.Intn(100))
	courses=append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("update one Course ")
	w.Header().Set("Content-Type:","application/json")

	params:=mux.Vars(r)

	for index,course:= range courses{
         if course.CourseId==params["id"]{
			courses=append(courses[:index],courses[:index+1]... )
            var course Course
			_=json.NewDecoder(r.Body).Decode(&course)
			 course.CourseId=params["id"]
			 courses=append(courses, course)
			 json.NewEncoder(w).Encode(course) 
			 return
		 }	
	}
}	

func deleteAllcourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Delete all the course you have")
	w.Header().Set("Content-Type:","application/json")

	params:=mux.Vars(r)

	for index,course:=range courses{
		if course.CourseId==params["id"] {
			courses=append(courses[:index], courses[index+1:]...)
            break
		}
		
	}


}