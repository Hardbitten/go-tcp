package handlers

import (
	"main/models"
	"main/opcodes"
	"main/serializers"
)

func HandlePlayerUpdatePosition(data *serializers.ByteBuffer, player *models.Player) {

	bf := serializers.NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_UPDATE_POSITION)

	bf.WriteUInt32(player.ID)

	bf.WriteFloat(data.ReadFloat()) // p X
	bf.WriteFloat(data.ReadFloat()) // p Y
	bf.WriteFloat(data.ReadFloat()) // p Z

	player.BroadcastLobby(bf)

}
