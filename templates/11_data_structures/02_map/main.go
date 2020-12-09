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
	names := map[string]string{
		"John":   "Smith",
		"James":  "Bond",
		"Jackie": "Chan",
		"Brad":   "Pitt",
	}

	for key, val := range names {
		log.Println(`` + key + ` - ` + val + ``)
	}

	err := textTemplate.ExecuteTemplate(os.Stdout, "tpl.gohtml", names)
	if err != nil {
		log.Fatal(err)
	}
}
