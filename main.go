package main

import (
	"log"
	"net/http"
	"html/template"
)

// User defenition
type User struct {
	Name string
	Age uint16
	Degree string
	WorkExperience float64
	Confidence float64
	Hobbies []string
}

func main() {
	handler()
}

func home(w http.ResponseWriter, r *http.Request) {
	user := User {
		"Ivan",
		18,
		"Junior Web-Developer", 
		0.0,
		4.5,
		[]string{
			"code", "skate", "ski",
			"walk around the city", "read novels",
		},
	}
	t, e := template.ParseFiles("templates/home.html")
	if e != nil {
		panic(e)
	}
	t.Execute(w, user)
}

func contact(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("templates/contact.html")
	if e != nil {
		panic(e)
	}
	t.Execute(w, contact)
}

func about(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("templates/about.html")
	if e != nil {
		panic(e)
	}
	t.Execute(w, about)
}

func handler() {
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about", about)
	log.Fatal(http.ListenAndServe(":8080", nil))
}