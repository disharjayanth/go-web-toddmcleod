package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var textTemplate *template.Template

// Person struct
type Person struct {
	Name string
	Age  int
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	textTemplate = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	person1 := Person{
		Name: "John",
		Age:  29,
	}

	person2 := Person{
		Name: "James",
		Age:  22,
	}

	people := []Person{person1, person2}

	err := textTemplate.ExecuteTemplate(os.Stdout, "tpl.gohtml", people)
	if err != nil {
		log.Fatal(err)
	}
}
