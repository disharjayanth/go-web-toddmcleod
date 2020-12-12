package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", formHandler)
	http.ListenAndServe(":3000", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	var s string

	if r.Method == http.MethodPost {
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Cannot Read the file uploaded", 500)
			return
		}
		defer file.Close()

		fmt.Println("\nFile:", file, "\nFile Header:", fileHeader, "\nErr:", err)

		sliceByteStr, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, "Cannot Read the file:", 500)
			return
		}

		s = string(sliceByteStr)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(w, "index.gohtml", s)
}
