package handlers

import (
	"main/models"
	"main/opcodes"
	"main/serializers"
)

func HandlePlayerUpdateRotation(data *serializers.ByteBuffer, player *models.Player) {

	bf := serializers.NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_UPDATE_ROTATION)

	bf.WriteUInt32(player.ID)

	// Y Axies
	y := data.ReadFloat()

	bf.WriteFloat(y)

	player.Rotation = y

	player.BroadcastLobby(bf)
}
