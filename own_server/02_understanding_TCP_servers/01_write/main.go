package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// TELNET IS USED TO RUN ON TCP NETWORK INORDER TO MAKE A REQUEST

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error Listening at PORT: 8080", err)
	}

	defer li.Close()

	for {
		Conn, err := li.Accept()
		if err != nil {
			log.Fatal("Error Accepting Connections:", err)
			return
		}

		io.WriteString(Conn, `\n Hello from TCP server.\n`)
		fmt.Fprintln(Conn, "How is your day?")
		fmt.Fprintf(Conn, "%v", "Well my day was good.")

		Conn.Close()
	}
}
