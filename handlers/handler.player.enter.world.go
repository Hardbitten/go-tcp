package handlers

import (
	"fmt"
	"main/models"
	"main/opcodes"
	"main/serializers"
)

func HandlePlayerEnterWorld(data *serializers.ByteBuffer, player *models.Player) {

	bf := serializers.NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_ENTER_WORLD)

	bf.WriteUInt32(player.ID)

	bf.WriteFloat(0)
	bf.WriteFloat(0)
	bf.WriteFloat(0)

	player.BroadcastLobby(bf)
	fmt.Printf("player [%d] joined to Game\n", player.ID)

	// send sync
	for _, lobbyPlayer := range player.Lobby.Players {
		if lobbyPlayer.ID != player.ID {

			// Get latest Data for user
			bfSync := serializers.NewByteBuffer()
			bfSync.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_SYNC)
			bfSync.WriteUInt32(lobbyPlayer.ID)

			x, y, z := lobbyPlayer.Position.X, lobbyPlayer.Position.Y, lobbyPlayer.Position.Z
			rotation := lobbyPlayer.Rotation
			bfSync.WriteFloat(x)
			bfSync.WriteFloat(y)
			bfSync.WriteFloat(z)
			bfSync.WriteFloat(rotation)
			player.Session.Conn.Write(bfSync.GetData())
		}

	}
}
