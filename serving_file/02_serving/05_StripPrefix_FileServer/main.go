package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/assests/toby.jpg">`)
}

func main() {
	http.HandleFunc("/dog", dog)
	http.Handle("/assests/", http.StripPrefix("/assests", http.FileServer(http.Dir("./assests"))))
	http.ListenAndServe(":3000", nil)
}
