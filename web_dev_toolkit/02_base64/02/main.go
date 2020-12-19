package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "Hello there!"

	base64EncodedString := base64.StdEncoding.EncodeToString([]byte(str))

	fmt.Println("Orginal String:", str)
	fmt.Println("Base 64 Standard Encoded String:", base64EncodedString)
	fmt.Println("Len of orginal string:", len(str))
	fmt.Println("Len of base64 encoded string:", len(base64EncodedString))
}
