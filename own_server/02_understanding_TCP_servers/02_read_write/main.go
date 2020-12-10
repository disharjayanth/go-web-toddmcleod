package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// telnet localhost 8080 in another terminal and then, type something and check server termianl, you can see what u sent

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to Listen from PORT:8080", err)
	}
	defer li.Close()

	for {
		Conn, err := li.Accept()
		if err != nil {
			log.Fatal("Failed to listen to connection:", err)
		}
		go handle(Conn)
	}
}

func handle(Conn net.Conn) {
	scanner := bufio.NewScanner(Conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(Conn, "I heard what you said: %v", ln)
	}
	defer Conn.Close()

	// Code never gets till here since it keeps listening.
	fmt.Println("handle func ended.")
}
