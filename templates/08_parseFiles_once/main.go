package main

import (
	"log"
	"os"
	"text/template"
)

var htmlTemps *template.Template

func init() {
	htmlTemps = template.Must(template.ParseGlob(`templates/*`))
}

func main() {
	err := htmlTemps.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
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
