package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	// add to serve pic
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r)
	if r.Method == http.MethodPost {
		upFile, upHeader, err := r.FormFile("uploadPicFile")
		if err != nil {
			fmt.Println("Cannot get file from form:", err)
		}
		defer upFile.Close()

		// Get the extension of uploaded file
		fileExtension := strings.Split(upHeader.Filename, ".")[1]

		// Create sha for file content
		shaHash := sha1.New()

		// Copy file content to hash, which creates hash which is unique to any given value/content
		io.Copy(shaHash, upFile)

		// hash.Sum([]byte) []byte returns slice of byte
		fname := fmt.Sprintf("%x", shaHash.Sum(nil)) + "." + fileExtension

		currentWDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Cannot get pwd:", err)
		}

		filePath := filepath.Join(currentWDir, "public", "pics", fname)

		newFile, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file at path public/pics:", err)
		}
		defer newFile.Close()

		// offset at 0 meaning start of file, 1 current, 2 end and whence at 0
		upFile.Seek(0, 0)

		io.Copy(newFile, upFile)

		cookie = appendValue(w, cookie, fname)
	}

	sliceCookieValue := strings.Split(cookie.Value, "|")

	tpl.ExecuteTemplate(w, "index.html", sliceCookieValue[1:])
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("userLogin")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "userLogin",
			Value: uuid.NewV4().String(),
		}
		http.SetCookie(w, cookie)
	}
	return cookie
}

func appendValue(w http.ResponseWriter, cookie *http.Cookie, fname string) *http.Cookie {
	stringCookieValue := cookie.Value

	if !strings.Contains(stringCookieValue, fname) {
		stringCookieValue += "|" + fname
	}

	cookie.Value = stringCookieValue

	http.SetCookie(w, cookie)

	return cookie
}
