package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main() {
	// Define the server address and port
	serverAddress := "35.212.209.66:4222" // Replace with your server's address and port

	// Connect to the TCP server
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close() // Ensure the connection is closed when the function exits

	fmt.Printf("Connected to TCP server at %s\n", serverAddress)

	// Create a reader to get input from the console
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message to send (type 'exit' to quit): ")
		input, _ := reader.ReadString('\n')
		messageToSend := strings.TrimSpace(input)

		if messageToSend == "exit" {
			fmt.Println("Exiting TCP client.")
			return
		}

		// Send the message to the server
		_, err = conn.Write([]byte(messageToSend + "\n")) // Add newline for server to read as line
		if err != nil {
			fmt.Printf("Error sending data: %v\n", err)
			return
		}

		// Read the response from the server
		serverReader := bufio.NewReader(conn)
		response, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading response from server: %v\n", err)
			return
		}
		fmt.Printf("Received from server: %s", response)
	}
}
