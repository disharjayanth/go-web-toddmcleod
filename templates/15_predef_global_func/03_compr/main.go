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
	scores := struct {
		Score1 int
		Score2 int
	}{
		Score1: 7,
		Score2: 9,
	}

	err := textTemplate.ExecuteTemplate(os.Stdout, "tpl.gohtml", scores)
	if err != nil {
		log.Fatal(err)
	}
}
