package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Connection address
	serverAddr := "localhost:8080"

	// Connect to server
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Server connectio error: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("Peer B: Connected to server.")

	// Open file for saving
	file, err := os.Create("received_file.txt")
	if err != nil {
		fmt.Println("Creating file failure: ", err)
		return
	}
	defer file.Close()

	// Read file from connection and write it to file
	_, err = io.Copy(file, conn)
	if err != nil {
		fmt.Println("Receving file failure: ", err)
	}

	fmt.Println("Peer B: File received and saved.")
}
