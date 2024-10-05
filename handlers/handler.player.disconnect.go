package handlers

import (
	"fmt"
	"main/models"
	"main/serializers"
)

func HandlePlayerDisconnect(data *serializers.ByteBuffer, player *models.Player) {
	bf := serializers.SerializePlayerDisconnect(player.ID)

	player.BroadcastLobby(bf)
	fmt.Printf("player [%d] Disconnected from Game\n", player.ID)
}
