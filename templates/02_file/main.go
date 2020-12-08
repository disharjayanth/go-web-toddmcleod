package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := `Todd McLoed`

	tpl := `
	<html>
	<head>
	<meta charset="UTF-8">
	<title>Go Web Programming</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	_, err = io.Copy(file, strings.NewReader(tpl))
	if err != nil {
		log.Fatal("Error writing to the file:", err)
	}

	log.Println("Exited.")
}
