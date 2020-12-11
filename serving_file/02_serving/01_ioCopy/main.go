package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
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
		http.Error(w, "Cannot read image file", http.StatusNotFound)
	}
	defer imgFile.Close()

	io.Copy(w, imgFile)
}
