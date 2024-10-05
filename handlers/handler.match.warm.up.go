package handlers

import (
	"fmt"
	m "main/matchmaking"
	"main/models"
	"main/utils"
)

func HandleMatchWarmUp(data *utils.ByteBuffer, player *models.Player) {
	fmt.Println("Handling Match Warm Up")
	m.AddPlayerToMatchmaking(player, true)
}
