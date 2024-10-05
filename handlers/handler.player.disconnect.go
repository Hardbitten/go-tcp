package handlers

import (
	"fmt"
	"main/models"
	"main/serializers"
	"main/utils"
)

func HandlePlayerDisconnect(data *utils.ByteBuffer, player *models.Player) {
	bf := serializers.SerializePlayerDisconnect(player.ID)

	player.BroadcastLobby(bf)
	fmt.Printf("player [%d] Disconnected from Game\n", player.ID)
}
