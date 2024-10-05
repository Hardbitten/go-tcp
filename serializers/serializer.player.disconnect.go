package serializers

import (
	"main/opcodes"
)

func SerializePlayerDisconnect(PlayerID uint32) *ByteBuffer {
	bf := NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_DISCONNECT)
	bf.WriteUInt32(PlayerID)

	// Remove Player from Lobby and other stored stuff...
	return bf
}
