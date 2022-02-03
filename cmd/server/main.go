package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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
	if err := handler(); err != nil {
		log.Fatal(err)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	user := User {
		"Ivan",
		18,
		"Self-educated Backend-Developer", 
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

func handler() error {
	http.HandleFunc("/", home)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about", about)

	fmt.Println("Starting server on URL: `localhost:5000`")
	
	return http.ListenAndServe(":5000", nil)
}