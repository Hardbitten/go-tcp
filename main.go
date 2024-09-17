package main

import (
	"fmt"
	"main/models"
	"main/serializers"
	"net"
)

func main() {
	// Start server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started on :8080")

	// Main loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Create session for the connected player
		session := models.NewSession(conn)
		go handleConnection(session)
	}
}

func handleConnection(session *models.Session) {
	defer session.Conn.Close()

	// Process incoming messages
	for {
		buffer := make([]byte, 1024)
		n, err := session.Conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		data := serializers.NewByteBufferWithData(buffer[:n])
		opcode := data.ReadInt16() // Assuming opcode is 16-bit

		// Call handler based on opcode
		if handler, ok := Handlers.MessageHandlers[opcode]; ok {
			handler(data, session)
		} else {
			fmt.Println("Unknown opcode:", opcode)
		}
	}
}
