package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println("command argument 0:", os.Args[0])
	fmt.Println("command argument 1:", os.Args[1])

	tpl := `
	<html>
	<head>
	<meta charset="UTF-8">
	<title>Go Web Programming with command args.</title>
	</head>
	<body>
	<h2>` + name + `</h2>
	</body>
	</html>
	`

	htmlFile, err := os.Create(`index.html`)
	if err != nil {
		log.Fatal("Error creating new file:", err)
	}
	defer htmlFile.Close()

	_, err = io.Copy(htmlFile, strings.NewReader(tpl))
	if err != nil {
		log.Fatal("Error copying contents to new file:", err)
	}

	log.Println("Exited.")
}
