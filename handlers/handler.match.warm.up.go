package handlers

import (
	"fmt"
	m "main/matchmaking"
	"main/models"
	"main/serializers"
)

func HandleMatchWarmUp(data *serializers.ByteBuffer, player *models.Player) {
	fmt.Println("Handling Match Warm Up")
	m.AddPlayerToMatchmaking(player, true)
}
