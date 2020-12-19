package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", loginPage)
	http.HandleFunc("/auth", authPage)

	http.ListenAndServe(":3000", nil)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "session",
			Value: "",
		}
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		cookie.Value = email + "|" + getHash(email)
	}

	http.SetCookie(w, cookie)

	io.WriteString(w, `
	<!DOCTYPE html>
	<html>
		<body>
			<form method="POST" action="/">
				<label>Email:</label>
				<input name="email" type="email />
				<input type="submit" />
			</form>
			<a href="/auth">Validate this :`+cookie.Value+`</a>
		</body>
	</html>
	`)
}

func authPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if cookie.Value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	sliceOfStringCookieValue := strings.Split(cookie.Value, "|")
	email := sliceOfStringCookieValue[0]
	hash := sliceOfStringCookieValue[1]

	hashCheck := getHash(email)

	if hash != hashCheck {
		fmt.Println("Hash did not match.")
		fmt.Println("User hash from browser:", hash)
		fmt.Println("Recreated hash from email:", hashCheck)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	io.WriteString(w, `
	<!DOCTYPE html>
	<html>
		<body>
			<h1>Received hash:`+hash+`</h1>
			<h1>Verified hash:`+hashCheck+`</h1>
		</body>
	</html>
	`)
}

func getHash(str string) string {
	hash := hmac.New(sha256.New, []byte("privateKey"))
	io.WriteString(hash, str)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
