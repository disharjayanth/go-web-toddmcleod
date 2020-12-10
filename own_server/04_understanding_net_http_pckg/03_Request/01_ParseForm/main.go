package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(w, "template.gohtml", r.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	var h hotdog
	http.ListenAndServe(":3000", h)
}
