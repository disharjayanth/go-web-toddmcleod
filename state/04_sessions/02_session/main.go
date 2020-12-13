package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// User struct
type User struct {
	Email     string
	FirstName string
	LastName  string
}

var tpl *template.Template
var dbUsers = make(map[string]User)
var dbSessions = make(map[string]string)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", getFormOrSubmitForm)
	http.HandleFunc("/user", getUserDetails)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":3000", nil)
}

func getFormOrSubmitForm(w http.ResponseWriter, r *http.Request) {
	// inital load show form and show the user details
	cookie, err := r.Cookie("session")
	if err != nil {
		// if cookie is not present in request then create one
		sID := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
	}

	var u User
	if userEmail, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[userEmail]
	}

	if r.Method == http.MethodPost {
		useremail := r.FormValue("useremail")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		u = User{
			Email:     useremail,
			FirstName: fname,
			LastName:  lname,
		}
		dbSessions[cookie.Value] = useremail
		dbUsers[useremail] = u
	}

	tpl.ExecuteTemplate(w, "index.html", u)
}

func getUserDetails(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	userEmail, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	user := dbUsers[userEmail]
	tpl.ExecuteTemplate(w, "bar.html", user)
}
