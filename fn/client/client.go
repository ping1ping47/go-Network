package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("----------------------------")
	fmt.Println("\x1b[1;34mSimple Chat Client\x1b[0m")
	fmt.Println("----------------------------")

	fmt.Print("\x1b[36mEnter username: \x1b[0m")
	var username string
	fmt.Scanln(&username)

	fmt.Print("\x1b[32mEnter password: \x1b[0m")
	var password string
	fmt.Scanln(&password)

	sendData(username, password)
}

func sendData(username, password string) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	data := username + " " + password
	fmt.Fprintf(conn, data)

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Print("\x1b[31mServer response: ", response, "\x1b[0m\n")
}
