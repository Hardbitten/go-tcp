package deserializers

import "main/utils"

func DeserializePlayerMessageChat(data *utils.ByteBuffer) (uint16, string, error) {
	// Read the length of the string message
	messageLength := data.ReadUInt16()

	// Read the actual message string
	message := data.ReadString(uint(messageLength))

	return messageLength, message, nil
}
