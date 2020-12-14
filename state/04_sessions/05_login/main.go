package main

import (
	"html/template"
	"log"
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
	bsPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	if err != nil {
		log.Fatalln("Error initialising password.")
		return
	}
	dbUsers["example@example.com"] = User{
		Email:     "example@example.com",
		Password:  bsPassword,
		FirstName: "example",
		LastName:  "example",
	}
}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/user", userDetailsPage)
	http.HandleFunc("/signup", signupPage)
	http.HandleFunc("/signin", signinPage)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":3000", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	user := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.html", user)
}

func userDetailsPage(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	user := getUser(w, r)
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

		if _, ok := dbUsers[email]; ok {
			http.Error(w, "User already registered with given email.", http.StatusForbidden)
			return
		}

		// create a session and cookie
		cookie := &http.Cookie{
			Name:  "userLogin",
			Value: uuid.NewV4().String(),
		}

		http.SetCookie(w, cookie)

		dbSessions[cookie.Value] = email

		// store user in dbUsers
		// ency before it
		bsPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		user := User{
			Email:     email,
			Password:  bsPassword,
			FirstName: fname,
			LastName:  lname,
		}

		dbUsers[email] = user
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
			http.Error(w, "User email or password does not match.", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
		if err != nil {
			http.Error(w, "User email or password does not match.", http.StatusForbidden)
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
