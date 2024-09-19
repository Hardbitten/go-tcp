package handlers

import (
	"main/models"
	"main/serializers"
)

// Define a type for the handler function
type HandlerType func(data *serializers.ByteBuffer, player *models.Player)

var Handlers = map[uint16]HandlerType{
	0x01: HandleAuthLogin, // AUTH_LOGIN

	0x500: HandlePlayerEnterWorld,
	// 0x101: HandlePlayerEnterWorld1,
	// 0x102: HandlePlayerEnterWorld2,
}
