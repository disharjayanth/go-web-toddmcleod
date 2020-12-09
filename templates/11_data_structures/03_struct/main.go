package main

import (
	"os"
	"text/template"
)

var textTemp *template.Template

// Person struct
type Person struct {
	Name string
	Age  int
}

func init() {
	textTemp = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	john := Person{
		Name: "John",
		Age:  28,
	}
	textTemp.Execute(os.Stdout, john)
}
