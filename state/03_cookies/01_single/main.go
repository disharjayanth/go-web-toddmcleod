package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/read", readCookie)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":3000", nil)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some important value",
	})
	fmt.Fprintln(w, "Cookie is written your browser.")
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, "Error reading cookie in server", http.StatusInternalServerError)
	}
	fmt.Fprintln(w, "Your cookie:", cookie)
}
