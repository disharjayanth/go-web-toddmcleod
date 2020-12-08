package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	textTemplate, err := template.ParseFiles("one.txt")
	if err != nil {
		log.Fatal("Error parsing one.txt:", err)
	}

	err = textTemplate.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("Error execting parsed file:", err)
	}

	textTemplate, err = textTemplate.ParseFiles("two.txt", "three.txt", "four.txt")
	if err != nil {
		log.Fatal("Error parsing one.txt, two.txt, three.txt")
	}

	err = textTemplate.ExecuteTemplate(os.Stdout, "two.txt", nil)
	if err != nil {
		log.Fatal("Error executing two.txt:", err)
	}

	err = textTemplate.ExecuteTemplate(os.Stdout, "three.txt", nil)
	if err != nil {
		log.Fatal("Error executing three.txt:", err)
	}

	err = textTemplate.ExecuteTemplate(os.Stdout, "four.txt", nil)
	if err != nil {
		log.Fatal("Error executing four.txt:", err)
	}

	err = textTemplate.ExecuteTemplate(os.Stdout, "one.txt", nil)
	if err != nil {
		log.Fatal("Error executing one.txt:", err)
	}

	log.Println("Exited.")
}
