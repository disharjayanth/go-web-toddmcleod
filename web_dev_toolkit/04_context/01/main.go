package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainPage)

	http.ListenAndServe(":3000", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
