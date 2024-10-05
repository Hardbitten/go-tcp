package serializers

import (
	"main/opcodes"
)

func SerializePlayerRotation(PlayerID uint32, Rotation float32) *ByteBuffer {
	bf := NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_UPDATE_ROTATION)
	bf.WriteUInt32(PlayerID)
	bf.WriteFloat(Rotation)
	return bf
}
