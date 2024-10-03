package handlers

import (
	"fmt"
	"main/models"
	"main/opcodes"
	"main/serializers"
)

func HandlePlayerDisconnect(data *serializers.ByteBuffer, player *models.Player) {
	bf := serializers.NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_DISCONNECT)
	bf.WriteUInt32(player.ID)

	player.Lobby.RemovePlayer(player.ID)

	player.BroadcastLobby(bf)
	fmt.Printf("player [%d] Disconnected from Game\n", player.ID)
}
