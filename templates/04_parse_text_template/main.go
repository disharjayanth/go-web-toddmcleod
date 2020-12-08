package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	htmlTemplate, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal("Error parsing the .gohtml file which contains html:", err)
	}

	err = htmlTemplate.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Error executing parsed files:", err)
	}
}
