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

func main() {
	xs := []string{"one", "two", "three", "four"}
	data := struct {
		Words []string
		Lname string
	}{
		Words: xs,
		Lname: "Smith",
	}

	err := textTemplate.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}
