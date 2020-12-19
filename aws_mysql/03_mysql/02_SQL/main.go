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
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/readall", readall)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err = http.ListenAndServe(":3000", nil)
	check(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "At main page.")
}

func readall(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT name FROM person;")
	check(err)

	var s, name string
	s = "Retrieved Rows are:\n"

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}

	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, r *http.Request) {
	preparedStmt, err := db.Prepare("CREATE TABLE customer (name VARCHAR(255));")
	check(err)

	defer preparedStmt.Close()

	rows, err := preparedStmt.Exec()
	check(err)

	n, err := rows.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Create Table customer:", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
	preparedStmt, err := db.Prepare(`INSERT INTO customer VALUES ("Coleman");`)
	check(err)

	defer preparedStmt.Close()

	rows, err := preparedStmt.Exec()
	check(err)

	n, err := rows.RowsAffected()
	check(err)

	fmt.Fprintln(w, "No. of rows affected:", n)
}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM customer;")
	check(err)

	var names []string
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)

		names = append(names, name)
	}

	fmt.Fprintln(w, names)
}

func update(w http.ResponseWriter, r *http.Request) {
	preparedStmt, err := db.Prepare(`UPDATE customer SET name="Jimmy" WHERE name="James";`)
	check(err)

	rows, err := preparedStmt.Exec()
	check(err)

	n, err := rows.RowsAffected()
	check(err)

	fmt.Fprintln(w, "No. of rows affected:", n)
}

func delete(w http.ResponseWriter, r *http.Request) {
	preparedStmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jimmy";`)
	check(err)

	rows, err := preparedStmt.Exec()
	check(err)

	n, err := rows.RowsAffected()
	check(err)

	fmt.Fprintln(w, "No. of rows affected:", n)
}

func drop(w http.ResponseWriter, r *http.Request) {
	preparedStmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)

	_, err = preparedStmt.Exec()
	check(err)

	fmt.Fprintln(w, "DROPPED TABLE customer.")
}

func check(err error) {
	if err != nil {
		fmt.Println("Error connecting database:", err)
		return
	}
}
