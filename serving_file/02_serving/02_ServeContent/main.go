package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dogpic", dogPic)
	http.ListenAndServe(":3000", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	imgFile, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "img file cannot me opened", http.StatusNotFound)
	}
	defer imgFile.Close()

	info, err := imgFile.Stat()
	if err != nil {
		http.Error(w, "cannot find img info", 404)
	}

	http.ServeContent(w, r, imgFile.Name(), info.ModTime(), imgFile)
}
