package main

import (
	"io"
	"net/http"
)

func urlData(w http.ResponseWriter, r *http.Request) {
	valueString := r.FormValue("key1")
	io.WriteString(w, `The value of url query string: `+valueString+`.`)
}

func main() {
	http.HandleFunc("/", urlData)
	http.ListenAndServe(":3000", nil)
}
