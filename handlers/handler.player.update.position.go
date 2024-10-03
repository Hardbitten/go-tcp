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

	x, y, z := data.ReadFloat(), data.ReadFloat(), data.ReadFloat()

	bf.WriteFloat(x) // p X
	bf.WriteFloat(y) // p Y
	bf.WriteFloat(z) // p Z

	player.Position = models.NewVector(x, y, z)

	player.BroadcastLobby(bf)

}
