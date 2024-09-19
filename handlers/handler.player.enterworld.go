package handlers

import (
	"fmt"
	"main/models"
	"main/serializers"
)

func HandlePlayerEnterWorld(data *serializers.ByteBuffer, player *models.Player) {

	fmt.Println("new Player entered to the world.")
	// player.Session.Broadcast()
}
