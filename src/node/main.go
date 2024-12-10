package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	nodeName := os.Getenv("NODE_NAME")
	fmt.Printf("Starting %s...\n", nodeName)

	// Listen on a TCP port
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error starting server for %s: %v\n", nodeName, err)
		return
	}
	defer listener.Close()

	fmt.Printf("%s is listening on port 8080\n", nodeName)

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection for %s: %v\n", nodeName, err)
			continue
		}
		go handleConnection(conn, nodeName)
	}
}

func handleConnection(conn net.Conn, nodeName string) {
	defer conn.Close()

	// Read message
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("Error reading from connection in %s: %v\n", nodeName, err)
		return
	}

	message := string(buf[:n])
	fmt.Printf("%s received: %s\n", nodeName, message)

	// Respond to the sender
	_, err = conn.Write([]byte(fmt.Sprintf("Message received by %s", nodeName)))
	if err != nil {
		fmt.Printf("Error writing to connection in %s: %v\n", nodeName, err)
		return
	}
}
