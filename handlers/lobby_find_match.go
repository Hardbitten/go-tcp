package handlers

import (
	"fmt"
	"main/models"
	"main/serializers"
	// Import other necessary packages
)

func HandleLobbyFindMatch(data *serializers.ByteBuffer, player *models.Player) {
	// Deserialize data
	// Process finding match
	fmt.Println("Handling LOBY_FIND_MATCH")
}
