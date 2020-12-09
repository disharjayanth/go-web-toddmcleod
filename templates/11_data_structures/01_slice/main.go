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
	names := []string{"John", "Smith", "Brad", "Stephen", "Alex"}
	err := textTemplate.ExecuteTemplate(os.Stdout, "tpl.gohtml", names)
	if err != nil {
		log.Fatal(err)
	}
}
