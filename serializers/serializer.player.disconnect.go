package serializers

import (
	"main/opcodes"
	"main/utils"
)

func SerializePlayerDisconnect(PlayerID uint32) *utils.ByteBuffer {
	bf := utils.NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_DISCONNECT)
	bf.WriteUInt32(PlayerID)

	// Remove Player from Lobby and other stored stuff...
	return bf
}
