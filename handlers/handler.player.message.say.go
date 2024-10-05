package handlers

import (
	"fmt"
	"main/models"
	"main/serializers"
)

func HandlePlayerMessageChat(data *serializers.ByteBuffer, player *models.Player) {

	// Read the length of the string message
	messageLength := data.ReadUInt16()

	// Read the actual message string
	message := data.ReadString(uint(messageLength))

	// Prepare the response buffer to broadcast the message
	bf := serializers.SerializePlayerMessageChat(player.ID, messageLength, message)

	fmt.Printf("player %d says %s\n", player.ID, message)
	// Broadcast the message to all players in the lobby
	player.BroadcastLobby(bf)
}
