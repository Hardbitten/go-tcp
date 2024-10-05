package handlers

import (
	"fmt"
	"main/models"
	"main/serializers"
	"main/utils"
)

func HandlePlayerEnterWorld(data *utils.ByteBuffer, player *models.Player) {

	bf := serializers.SerializePlayerEnterWorld(player.ID, player.Position, player.Rotation, player.Name)

	player.BroadcastLobby(bf)
	fmt.Printf("player [%d] joined to Game\n", player.ID)

	// send sync
	for _, lobbyPlayer := range player.Lobby.Players {
		if lobbyPlayer.ID != player.ID {
			bfSync := serializers.SerializePlayerSync(lobbyPlayer.ID, lobbyPlayer.Position, lobbyPlayer.Rotation, lobbyPlayer.Name)
			player.Session.Conn.Write(bfSync.GetData())
		}

	}
}
