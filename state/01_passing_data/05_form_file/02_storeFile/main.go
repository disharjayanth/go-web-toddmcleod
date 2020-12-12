package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", formData)
	http.ListenAndServe(":3000", nil)
}

func formData(w http.ResponseWriter, r *http.Request) {
	var s string

	if r.Method == http.MethodPost {
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Cannot open file at server:", 500)
			return
		}
		defer file.Close()

		bs, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, "Cannot read file at server:", 500)
			return
		}

		s = string(bs)

		newFile, err := os.Create(filepath.Join("./users/", fileHeader.Filename))
		if err != nil {
			http.Error(w, "Cannot store the given file:", 500)
			return
		}
		defer newFile.Close()

		_, err = newFile.Write(bs)
		if err != nil {
			http.Error(w, `Cannot write to newly created file: `+err.Error()+``, 500)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(w, "index.gohtml", s)
}
