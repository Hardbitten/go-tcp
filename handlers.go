package main

import (
	handler "main/handlers"
	models "main/models"
	serializers "main/serializers"
	"net"
)

// Define a type for the handler function
type MessageHandler func(data *serializers.ByteBuffer, session *models.Session)

// Map of OPCODE to handler functions
var Handlers = map[uint16]MessageHandler{
	0x01: handler.HandleAuthLogin,          // AUTH_LOGIN
	0x02: handler.HandleLobbyFindMatch,     // LOBY_FIND_MATCH
	0x03: handler.HandleGameUpdatePosition, // GAME_UPDATE_POSITION
	// Add other handlers here
}

// Adjusted handleMessage function
func HandleMessage(opcode uint16, data []byte, conn net.Conn) {
	// Convert []byte to *ByteBuffer
	byteBuffer := serializers.NewByteBufferWithData(data)

	// Get the appropriate handler
	handler, exists := Handlers[opcode]
	if exists {
		handler(byteBuffer, conn)
	} else {
		// Handle unknown opcode
	}
}
