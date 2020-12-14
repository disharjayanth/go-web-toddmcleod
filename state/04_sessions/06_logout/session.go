package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) User {
	cookie, err := r.Cookie("userLogin")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "userLogin",
			Value: uuid.NewV4().String(),
		}
	}

	http.SetCookie(w, cookie)

	var user User

	if userEmail, ok := dbSessions[cookie.Value]; ok {
		user = dbUsers[userEmail]
	}

	return user
}

func alreadyLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("userLogin")
	if err != nil {
		return false
	}

	userEmail := dbSessions[cookie.Value]
	_, ok := dbUsers[userEmail]
	return ok
}
