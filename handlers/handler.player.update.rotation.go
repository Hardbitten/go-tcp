package handlers

import (
	"main/models"
	"main/serializers"
)

func HandlePlayerUpdateRotation(data *serializers.ByteBuffer, player *models.Player) {
	// Y Axies
	y := data.ReadFloat()
	bf := serializers.SerializePlayerRotation(player.ID, y)

	player.Rotation = y

	player.BroadcastLobby(bf)
}
