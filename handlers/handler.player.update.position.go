package handlers

import (
	"main/models"
	"main/serializers"
	utils "main/utils"
)

func HandlePlayerUpdatePosition(data *serializers.ByteBuffer, player *models.Player) {

	position := utils.NewVector(data.ReadFloat(), data.ReadFloat(), data.ReadFloat())
	bf := serializers.SerializePlayerPosition(player.ID, position)

	player.Position = position

	player.BroadcastLobby(bf)

}
