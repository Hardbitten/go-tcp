package handlers

import (
	"fmt"
	matchmaking "main/matchmaking"
	"main/models"
	"main/utils"
)

func HandleMatchFindStart(data *utils.ByteBuffer, player *models.Player) {
	fmt.Println("Handling MATCH_FIND_START")
	isWarmup := data.ReadBool() // Example: read from request whether it's a warm-up lobby or not
	matchmaking.AddPlayerToMatchmaking(player, isWarmup)
}
