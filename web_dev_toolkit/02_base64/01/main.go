package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "Hi there!"

	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	base64EncodedString := base64.NewEncoding(encodeStd).EncodeToString([]byte(str))

	fmt.Println("Length of orginal String:", len(str))
	fmt.Println("Length of base64 encoded string:", len(base64EncodedString))

	fmt.Println("Orginal string:", str)
	fmt.Println("Base 64 encoded string:", base64EncodedString)
}
