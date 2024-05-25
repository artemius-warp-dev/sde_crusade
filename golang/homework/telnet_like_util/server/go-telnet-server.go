package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Welcome to the Go Telnet Server!\n"))

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Received: ", message)
		if message == "quit" || message == "exit" {
			break
		}
		conn.Write([]byte("You said: " + message + "\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from connection: ", err)
		return
	}
}

func main() {
	address := "localhost:8080"

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server is listening on", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		fmt.Println("New connection from: ", conn.RemoteAddr())

		go handleConnection(conn)
	}
}
