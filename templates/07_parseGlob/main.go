package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	htmlTemps, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatal("Error parsing Glob of .gohtml files: ", err)
	}

	err = htmlTemps.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Error executing parsed file", err)
	}

	err = htmlTemps.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = htmlTemps.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = htmlTemps.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = htmlTemps.ExecuteTemplate(os.Stdout, "four.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
