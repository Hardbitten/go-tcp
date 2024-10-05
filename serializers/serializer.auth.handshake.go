package serializers

import (
	"main/enums"
	"main/opcodes"
	"main/utils"
)

// DeserializeAuthLogin deserializes data for AUTH_LOGIN opcode
func SerializeAuthHandshake(data *utils.ByteBuffer, status enums.AuthHandshakeStatus) *utils.ByteBuffer {
	bf := utils.NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_AUTH_HANDSHAKE)
	bf.WriteUInt16(uint16(status))
	return bf
}
