package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello....")
}

func main() {
	var d hotdog
	http.ListenAndServe(":3000", d)
}
