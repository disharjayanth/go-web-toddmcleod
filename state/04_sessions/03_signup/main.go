package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// User struct
type User struct {
	Email     string
	Password  string
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
	http.HandleFunc("/user", userDetails)
	http.HandleFunc("/signup", signUp)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":3000", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	user := getUser(r)
	tpl.ExecuteTemplate(w, "index.html", user)
}

func userDetails(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	user := getUser(r)
	tpl.ExecuteTemplate(w, "bar.html", user)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if r.Method == http.MethodPost {
		// get form values
		email := r.FormValue("email")
		password := r.FormValue("password")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")

		// user email already taken?
		if _, ok := dbUsers[email]; ok {
			http.Error(w, "User already registered with that email.", http.StatusForbidden)
			return
		}

		sID := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "userLogin",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)

		// store sid session id as key and value as email
		dbSessions[cookie.Value] = email

		user := User{
			Email:     email,
			Password:  password,
			FirstName: fname,
			LastName:  lname,
		}
		// store user with email as key
		dbUsers[email] = user

		// redirect to main page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.html", nil)
}
