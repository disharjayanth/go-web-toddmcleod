package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	str := "Hello World!"

	base64EncodedString := base64.StdEncoding.EncodeToString([]byte(str))

	fmt.Println("Orginal string:", str)
	fmt.Println("Base 64 encoded string:", base64EncodedString)

	sb, err := base64.StdEncoding.DecodeString(base64EncodedString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decoded string from base64:", string(sb))
}
