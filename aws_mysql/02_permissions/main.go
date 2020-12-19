package main

import (
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
	Role      string
}

var tpl *template.Template
var dbSessions = make(map[string]string)
var dbUsers = make(map[string]User)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/user", userDetailsPage)
	http.HandleFunc("/signup", signupPage)
	http.HandleFunc("/signin", signinPage)
	http.HandleFunc("/signout", signout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)
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

	if user.Role != "admin" {
		http.Error(w, "You need to be a admin to access this page.", http.StatusForbidden)
		return
	}

	tpl.ExecuteTemplate(w, "user.html", user)
}

func signupPage(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		role := r.FormValue("role")

		// check whether user already exists
		if _, ok := dbUsers[email]; ok {
			http.Error(w, "User registered with this email", http.StatusSeeOther)
			return
		}

		// create new cookie for user
		cookie := &http.Cookie{
			Name:  "userLogin",
			Value: uuid.NewV4().String(),
		}

		http.SetCookie(w, cookie)

		dbSessions[cookie.Value] = email

		// store new in dbUsers
		// encrypt 1st
		bsPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Error server", http.StatusInternalServerError)
			return
		}

		user := User{
			Email:     email,
			Password:  bsPassword,
			FirstName: fname,
			LastName:  lname,
			Role:      role,
		}

		dbUsers[email] = user

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func signinPage(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, ok := dbUsers[email]
		if !ok {
			http.Error(w, "Email or Password not correct.", http.StatusForbidden)
			return
		}

		// generate hash password
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
		if err != nil {
			http.Error(w, "Email or Password not correct.", http.StatusForbidden)
			return
		}

		cookie := &http.Cookie{
			Name:  "userLogin",
			Value: uuid.NewV4().String(),
		}

		http.SetCookie(w, cookie)

		dbSessions[cookie.Value] = email

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signin.html", nil)
}

func signout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusForbidden)
		return
	}

	cookie, _ := r.Cookie("userLogin")

	delete(dbSessions, cookie.Value)

	cookie = &http.Cookie{
		Name:   "userLogin",
		Value:  "",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
