package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Doggy....")
}

func cat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Cat....")
}

func main() {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)

	http.ListenAndServe(":3000", nil)
}
