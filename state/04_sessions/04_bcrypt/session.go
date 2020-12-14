package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) User {
	// get cookie
	c, err := r.Cookie("userLogin")
	if err != nil {
		c := &http.Cookie{
			Name:  "userLogin",
			Value: uuid.NewV4().String(),
		}
		http.SetCookie(w, c)
	}

	// if user exists
	var u User
	if userEmail, ok := dbSessions[c.Value]; ok {
		u = dbUsers[userEmail]
	}

	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("userLogin")
	if err != nil {
		return false
	}

	userEmail := dbSessions[c.Value]
	_, ok := dbUsers[userEmail]

	return ok
}
