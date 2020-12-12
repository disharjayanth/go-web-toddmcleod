package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

// Person struct
type Person struct {
	FirstName string
	LastName  string
	Age       string
	Subscribe bool
}

func formData(w http.ResponseWriter, r *http.Request) {
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	age := r.FormValue("age")
	subscribe := r.FormValue("subscribe") == "on"

	tpl.ExecuteTemplate(w, "index.gohtml", Person{
		FirstName: firstname,
		LastName:  lastname,
		Age:       age,
		Subscribe: subscribe,
	})
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", formData)
	http.ListenAndServe(":3000", nil)
}
