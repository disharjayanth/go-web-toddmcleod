package main

import (
	"log"
	"os"
	"text/template"
)

var textTemplate *template.Template

func init() {
	textTemplate = template.Must(template.ParseFiles("tpl.gohtml"))
}

// Person struct
type Person struct {
	Name    string
	Goodat  string
	IsAdmin bool
}

func main() {
	p1 := Person{
		Name:    "John",
		Goodat:  "Dancing",
		IsAdmin: true,
	}

	p2 := Person{
		Name:    "James",
		Goodat:  "Computers",
		IsAdmin: false,
	}

	p3 := Person{
		Name:    "",
		Goodat:  "Cooking",
		IsAdmin: true,
	}

	users := []Person{p1, p2, p3}

	err := textTemplate.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
	if err != nil {
		log.Fatal(err)
	}
}
