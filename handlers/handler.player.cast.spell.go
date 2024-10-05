package handlers

import (
	"fmt"
	"main/deserializers"
	"main/models"
	"main/serializers"
	"main/utils"
)

func HandlePlayerCastSpell(data *utils.ByteBuffer, player *models.Player) {
	SpellID, TargetPosition, err := deserializers.DeserializePlayerCastSpell(data)
	if err != nil {
		fmt.Printf("Error ! [%s]", err)
	}

	bf := serializers.SerializePlayerCastSpell(player.ID, SpellID, TargetPosition)
	player.BroadcastLobby(bf)
}
