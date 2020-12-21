package main

import (
	"encoding/json"
	"net/http"
)

// Person struct
type Person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/marshal", marshalPage)
	http.HandleFunc("/encode", encodePage)

	http.ListenAndServe(":3000", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	htmlString := `
	<html>
		<head>
		<meta charset="UTF-8">
		<title>Marshal and Encode</title>
		</head>
		<body>
		<h2><a href="/marshal">Marshal</a></h2>
		<h2><a href="/encode">Encode</a></h2>
		</body>
	</html>
	`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(htmlString))
}

func marshalPage(w http.ResponseWriter, r *http.Request) {
	person := Person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Spy", "British", "Action", "Movie"},
	}

	jsonByte, err := json.Marshal(person)
	if err != nil {
		http.Error(w, "Error marshalling struct type to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonByte)
}

func encodePage(w http.ResponseWriter, r *http.Request) {
	person := Person{
		Fname: "Miss Money",
		Lname: "Penny",
		Items: []string{"Actress"},
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		http.Error(w, "Error encoding struct type to JSON", http.StatusInternalServerError)
		return
	}
}
