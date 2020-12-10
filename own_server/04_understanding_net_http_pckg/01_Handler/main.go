package main

import (
	"fmt"
	"net/http"
)

// Handler is an interface with ServeHTTP method
// ServeHTTP method has signature of ServeHTTP(w http.ResponseWriter, r *http.Request)
// Any type with this method is also type of Handler interface.

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello....")
}

func main() {
	var d hotdog
	fmt.Println(d)
}
