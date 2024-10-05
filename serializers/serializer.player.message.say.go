package serializers

import (
	"main/opcodes"
)

func SerializePlayerMessageChat(PlayerID uint32, MessageLength uint16, Message string) *ByteBuffer {
	buffer := NewByteBuffer()

	buffer.WriteUInt16(opcodes.SMSG_MESSAGECHAT)
	buffer.WriteUInt16(opcodes.SMSG_MESSAGECHAT) // Server opcode for broadcasting the message
	buffer.WriteUInt32(PlayerID)                 // Player's ID who sent the message
	buffer.WriteUInt16(MessageLength)            // Include the message length
	buffer.WriteString(Message)                  // Write the actual message string

	return buffer
}
