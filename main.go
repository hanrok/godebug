package main

import (
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	_, err := io.Copy(conn, conn)
	if err != nil {
		log.Fatalf("Connection error %v", err)
	}
}

func main() {
	log.Println("Starting up repeater")
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Cannot listen on port 8080 %v", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Accepting connection error %v", err)
		}
		go handleConnection(conn)
	}
}
