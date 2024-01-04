package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("----------------------------")
	fmt.Println("\x1b[1;34mSimple Chat Server\x1b[0m")
	fmt.Println("----------------------------")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server is listening on :8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	request, err := readData(conn)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	if validateCredentials(request) {
		fmt.Println("\x1b[32mLogin successful!\x1b[0m")
		fmt.Fprintf(conn, "Hello\n")
	} else {
		fmt.Println("\x1b[31mLogin unsuccessful!\x1b[0m")
		fmt.Fprintf(conn, "Fail\n")
	}
}

func readData(conn net.Conn) (string, error) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}

func validateCredentials(request string) bool {
	credentials := strings.Split(request, " ")
	if len(credentials) != 2 {
		return false
	}

	username := strings.TrimSpace(credentials[0])
	password := strings.TrimSpace(credentials[1])

	return username == "std1" && password == "p@ssw0rd"
}
