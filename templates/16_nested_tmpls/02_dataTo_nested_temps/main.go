package main

import (
	"log"
	"os"
	"text/template"
)

var textTemplates *template.Template

func init() {
	textTemplates = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := textTemplates.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
	if err != nil {
		log.Fatal(err)
	}
}
