package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	Email     string
	Password  []byte
	FirstName string
	LastName  string
}

var tpl *template.Template
var dbSessions = make(map[string]string)
var dbUsers = make(map[string]User)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/signup", signUpPage)
	http.HandleFunc("/user", userDetailsPage)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":3000", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	user := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.html", user)
}

func userDetailsPage(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	user := getUser(w, r)
	tpl.ExecuteTemplate(w, "userDetails.html", user)
}

func signUpPage(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")

		if _, ok := dbUsers[email]; ok {
			http.Error(w, "Email already registed.", http.StatusForbidden)
			return
		}

		// create cookie for new user
		cookie := &http.Cookie{
			Name:  "userLogin",
			Value: uuid.NewV4().String(),
		}

		http.SetCookie(w, cookie)

		// store email value with sid key
		dbSessions[cookie.Value] = email

		// Before storing user in dbUsers , encrypt it first
		bsPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		user := User{
			Email:     email,
			Password:  bsPassword,
			FirstName: fname,
			LastName:  lname,
		}

		dbUsers[email] = user

		fmt.Println(dbSessions)
		fmt.Println(dbUsers)
		// After successfull registration and storing , redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.html", nil)
}
