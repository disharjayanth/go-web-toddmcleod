package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var textTemp *template.Template

func init() {
	textTemp = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := textTemp.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatal(`Error executing parse file :`, err)
	}
}
