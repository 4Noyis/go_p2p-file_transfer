// File sender peer

package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// File name and define port
	fileName := "sending_file.txt"
	port := ":8080"

	// Creating TCP listener
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Listening failure!", err)
		return
	}
	defer listener.Close()
	fmt.Println("Peer A: Waiting for connections...")

	// Accepting connection
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Connection accepting failure.", err)
		return
	}
	defer conn.Close()
	fmt.Println("Peer A: Connection Accepted.")

	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File openning failure.", err)
		return
	}
	defer file.Close()

	// Write file to connection
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("File sending failure.", err)
		return
	}

	fmt.Println("File sending complete.")

}
