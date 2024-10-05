package serializers

import (
	"main/opcodes"
	"main/utils"
)

func SerializePlayerMessageChat(PlayerID uint32, MessageLength uint16, Message string) *utils.ByteBuffer {
	buffer := utils.NewByteBuffer()
	buffer.WriteUInt16(opcodes.SMSG_OPCODE_MESSAGE_CHAT) // Server opcode for broadcasting the message
	buffer.WriteUInt32(PlayerID)                         // Player's ID who sent the message
	buffer.WriteUInt16(MessageLength)                    // Include the message length
	buffer.WriteString(Message)                          // Write the actual message string

	return buffer
}
