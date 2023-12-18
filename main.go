package main

import (
	"fmt"
	"log"
	"net"
)

const serverPort = "12345" // the same port as used in the client script

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error reading from client:", err)
		return
	}

	logEntry := string(buffer)
	fmt.Printf("Received log entry: %s\n", logEntry)
}

func main() {
	ln, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer ln.Close()

	fmt.Printf("Server listening on port %s\n", serverPort)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleClient(conn)
	}
}
