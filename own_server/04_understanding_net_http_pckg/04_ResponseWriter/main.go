package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(201)
	fmt.Fprintln(w, "Hello there! This is 1st example of responseWriter.")
}

func main() {
	var h hotdog
	http.ListenAndServe(":3000", h)
}
