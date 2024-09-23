package handlers

import (
	"fmt"
	matchmaking "main/matchmaking"
	"main/models"
	"main/serializers"
)

func HandleMatchFindStart(data *serializers.ByteBuffer, player *models.Player) {
	fmt.Println("Handling Match Find")
	matchmaking.AddPlayerToMatchmaking(player)
}
