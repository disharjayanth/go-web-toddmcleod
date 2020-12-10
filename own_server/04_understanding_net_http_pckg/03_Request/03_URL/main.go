package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("Error Parsing Form values:", err)
	}

	data := struct {
		Method string
		URL    *url.URL
		Values url.Values
	}{
		Method: r.Method,
		URL:    r.URL,
		Values: r.Form,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var h hotdog
	http.ListenAndServe(":3000", h)
}
