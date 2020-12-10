package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var textTemplate *template.Template

func init() {
	textTemplate = template.Must(template.ParseFiles("index.gohtml"))
}

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Method string
		Values url.Values
	}{
		Method: r.Method,
		Values: r.Form,
	}

	textTemplate.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	var h hotdog
	http.ListenAndServe(":3000", h)
}
