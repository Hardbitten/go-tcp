package serializers

import (
	"main/opcodes"
	"main/utils"
)

func SerializeMatchReady(PlayerID uint32) *utils.ByteBuffer {
	buffer := utils.NewByteBuffer()
	buffer.WriteUInt16(opcodes.SMSG_OPCODE_MATCH_READY)
	buffer.WriteUInt32(PlayerID)
	return buffer
}
