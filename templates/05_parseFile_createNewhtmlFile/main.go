package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	htmlTemplate, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal("Error parsing file:", err)
	}

	htmlFile, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}

	err = htmlTemplate.Execute(htmlFile, nil)
	if err != nil {
		log.Fatal("Error execting parsed file into new html file.")
	}
}
