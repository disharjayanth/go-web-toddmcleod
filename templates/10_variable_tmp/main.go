package main

import (
	"log"
	"os"
	"text/template"
)

var textTemp *template.Template

func init() {
	textTemp = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := textTemp.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatal(err)
	}
}
