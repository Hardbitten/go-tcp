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
	fmt.Printf("player [%d] joined to loby\n", player.ID)
}
