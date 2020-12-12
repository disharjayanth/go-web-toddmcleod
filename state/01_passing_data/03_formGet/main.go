package main

import (
	"html/template"
	"net/http"
)

func formData(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("q")
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		http.Error(w, "Cannot dispaly form", 404)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", value)
}

func main() {
	http.HandleFunc("/", formData)
	http.ListenAndServe(":3000", nil)
}
