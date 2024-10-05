package handlers

import (
	"fmt"
	"main/deserializers"
	"main/models"
	"main/serializers"
	"main/utils"
)

func HandlePlayerMessageChat(data *utils.ByteBuffer, player *models.Player) {

	messageLength, message, err := deserializers.DeserializePlayerMessageChat(data)
	if err != nil {
		fmt.Printf("Error ! [%s]", err)
	}

	// Prepare the response buffer to broadcast the message
	bf := serializers.SerializePlayerMessageChat(player.ID, messageLength, message)

	fmt.Printf("player %d says %s\n", player.ID, message)
	// Broadcast the message to all players in the lobby
	player.BroadcastLobby(bf)
}
