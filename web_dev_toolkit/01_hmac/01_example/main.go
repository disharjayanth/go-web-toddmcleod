package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	hmacHash := getCode("example@example.com")
	fmt.Println("Hmac hash:", hmacHash)
	hmacHash = getCode("example@example.com")
	fmt.Println("Hmac hash:", hmacHash)
}

// Given string is hashed to sha256
func getCode(s string) string {
	hmacHash := hmac.New(sha256.New, []byte("privateKey"))
	io.WriteString(hmacHash, s)
	return fmt.Sprintf("%x", hmacHash.Sum(nil))
}
