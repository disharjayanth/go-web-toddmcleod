package main

import (
	"io"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true ,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	io.WriteString(w, `The cookie stored in browser is: `+cookie.String()+``)
}
