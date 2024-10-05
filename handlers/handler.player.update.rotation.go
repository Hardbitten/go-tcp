package handlers

import (
	"fmt"
	"main/deserializers"
	"main/models"
	"main/serializers"
	"main/utils"
)

func HandlePlayerUpdateRotation(data *utils.ByteBuffer, player *models.Player) {

	rotation, err := deserializers.DeserializePlayerUpdateRotation(data)
	if err != nil {
		fmt.Printf("Error ! [%s]", err)
	}

	bf := serializers.SerializePlayerRotation(player.ID, rotation)
	player.Rotation = rotation

	player.BroadcastLobby(bf)
}
