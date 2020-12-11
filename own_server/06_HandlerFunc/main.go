package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog.")
}

func cat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Cat.")
}

func main() {
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/cat", http.HandlerFunc(cat))

	http.ListenAndServe(":3000", nil)
}
