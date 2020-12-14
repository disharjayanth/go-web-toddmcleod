package main

import "net/http"

func getUser(r *http.Request) User {
	var user User

	// get cookie
	c, err := r.Cookie("userLogin")
	if err != nil {
		return user
	}

	// if user exists already, get user
	if userEmail, ok := dbSessions[c.Value]; ok {
		user = dbUsers[userEmail]
	}

	return user
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("userLogin")
	if err != nil {
		return false
	}
	if userEmail, ok := dbSessions[c.Value]; ok {
		_, ok = dbUsers[userEmail]
		return ok
	}
	return false
}
