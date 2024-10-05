package serializers

import (
	"main/opcodes"
	"main/utils"
)

func SerializePlayerRotation(PlayerID uint32, Rotation float32) *utils.ByteBuffer {
	bf := utils.NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_UPDATE_ROTATION)
	bf.WriteUInt32(PlayerID)
	bf.WriteFloat(Rotation)
	return bf
}
