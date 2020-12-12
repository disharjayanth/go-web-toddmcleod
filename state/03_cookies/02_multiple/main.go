package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", setCookie1)
	http.HandleFunc("/read", readCookies)
	http.HandleFunc("/more", setCookie2)
	http.ListenAndServe(":3000", nil)
}

func setCookie1(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value of my-cookie",
	})

	fmt.Fprintf(w, "my-cookie written to browser cookie.")
}

func setCookie2(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "some value of general cookie",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some value of specific cookie",
	})

	fmt.Fprintln(w, "general and specific cookie written.")
}

func readCookies(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Value of my-cookie:", c1)
	}

	c2, err := r.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Value of general cookie:", c2)
	}

	c3, err := r.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Value of specific cookie:", c3)
	}
}
