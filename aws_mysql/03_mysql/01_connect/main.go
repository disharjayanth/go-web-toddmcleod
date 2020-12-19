package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "admin:password@tcp(database-1.czk8cw3v4g8o.ap-south-1.rds.amazonaws.com:3306)/test01?charset=utf8")
	check(err)

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Successfully connected to DB.")
}

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
