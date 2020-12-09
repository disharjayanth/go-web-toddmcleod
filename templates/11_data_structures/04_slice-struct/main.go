package main

import (
	"log"
	"os"
	"text/template"
)

var textTemplate *template.Template

// Person struct
type Person struct {
	Name    string
	Age     int
	Hobbie  string
	isAdmin bool
}

func init() {
	textTemplate = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	persons := []Person{
		Person{
			Name:    "John",
			Age:     28,
			Hobbie:  "singing",
			isAdmin: true,
		},
		Person{
			Name:    "James",
			Age:     30,
			Hobbie:  "football",
			isAdmin: false,
		},
		Person{
			Name:    "Alex",
			Age:     21,
			Hobbie:  "games",
			isAdmin: true,
		},
		Person{
			Name:    "Joe",
			Age:     22,
			Hobbie:  "golf",
			isAdmin: true,
		},
	}

	err := textTemplate.Execute(os.Stdout, persons)
	if err != nil {
		log.Fatal(err)
	}
}
