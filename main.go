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

	var player *models.Player = models.NewPlayer(session) // Initially, no player is assigned

	for {
		buffer := make([]byte, 1024)
		n, err := session.Conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		data := serializers.NewByteBufferWithData(buffer[:n])
		opcode := data.ReadUInt16() // Assuming opcode is 16-bit

		if opcode < 0x100 {
			// Call handler based on opcode
			if handler, ok := AuthHandlers[opcode]; ok {
				handler(data, session)
			} else {
				fmt.Println("Unknown opcode:", opcode)
			}
		} else {
			if player != nil {
				// Call handler based on opcode
				if handler, ok := GameHandlers[opcode]; ok {
					handler(data, player)
				} else {
					fmt.Println("Unknown opcode:", opcode)
				}
			}
		}

	}
}
