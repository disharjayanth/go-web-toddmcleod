package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "Bow Bow....")
	case "/cat":
		io.WriteString(w, "Meow Meow....")
	}
}

func main() {
	var h hotdog
	http.ListenAndServe(":3000", h)
}
