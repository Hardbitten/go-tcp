package main

import (
	"fmt"
	handlers "main/handlers"
	"main/models"
	"main/opcodes"
	"main/utils"
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
		// session :=
		player := models.NewPlayer(models.NewSession(conn))
		go handleConnection(player)
	}
}

func handleConnection(player *models.Player) {
	defer Disconnect(player)

	for {
		buffer := make([]byte, 1024)
		n, err := player.Session.Conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		data := utils.NewByteBufferWithData(buffer[:n])
		opcode := data.ReadUInt16() // Assuming opcode is 16-bit

		// Choose the correct handler map based on opcode
		// Call handler based on opcode
		if handler, ok := handlers.Handlers[opcode]; ok {
			handler(data, player)
		} else {
			fmt.Println("Unknown opcode:", opcode)
		}
	}
}

func Disconnect(player *models.Player) {

	bf := utils.NewByteBuffer()
	bf.WriteUInt16(opcodes.CMSG_OPCODE_PLAYER_DISCONNECT)
	handlers.HandlePlayerDisconnect(bf, player)

	player.Session.Conn.Close()
	fmt.Printf("Socket[%d] Disconnected\n", player.ID)
}
