package handlers

import (
	"fmt"
	"main/deserializers"
	"main/models"
	"main/serializers"
	"main/utils"
)

func HandlePlayerUpdatePosition(data *utils.ByteBuffer, player *models.Player) {

	position, err := deserializers.DeserializePlayerUpdatePosition(data)
	if err != nil {
		fmt.Printf("Error ! [%s]", err)
	}

	bf := serializers.SerializePlayerPosition(player.ID, position)

	player.Position = position

	player.BroadcastLobby(bf)

}
