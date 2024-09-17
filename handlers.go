package main

import (
	handler "main/handlers"
	models "main/models"
	serializers "main/serializers"
)

// Define a type for the handler function
type AuthHandler func(data *serializers.ByteBuffer, session *models.Session)
type GameHandler func(data *serializers.ByteBuffer, player *models.Player)

var AuthHandlers = map[uint16]AuthHandler{
	0x01: handler.HandleAuthLogin, // AUTH_LOGIN
	// Add other handlers here
}

var GameHandlers = map[uint16]GameHandler{
	0x02: handler.HandleLobbyFindMatch,     // LOBY_FIND_MATCH
	0x03: handler.HandleGameUpdatePosition, // GAME_UPDATE_POSITION
	// Add other handlers here
}
