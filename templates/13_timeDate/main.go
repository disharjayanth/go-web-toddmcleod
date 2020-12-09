package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var textTemplate *template.Template

var funcs = template.FuncMap{
	"formatDate": formatDate,
	"formatTime": formatTime,
}

func formatDate(dateTime time.Time) string {
	return dateTime.Format("Mon Jan 2 2006 03:04:05 IST")
}

func formatTime(timeTime time.Time) string {
	return timeTime.Format(time.Kitchen)
}

func init() {
	textTemplate = template.Must(template.New("").Funcs(funcs).ParseFiles("tpl.gohtml"))
}

func main() {
	err := textTemplate.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}
}
