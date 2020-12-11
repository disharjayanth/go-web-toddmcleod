package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="https://images.unsplash.com/photo-1566468161591-6ed35cb5565d?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=1950&q=80">`)
}

func main() {
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":3000", nil)
}
