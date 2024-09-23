package handlers

import (
	"fmt"
	"main/models"
	"main/opcodes"
	"main/serializers"
)

func HandlePlayerMessageSay(data *serializers.ByteBuffer, player *models.Player) {

	// Read the length of the string message
	messageLength := data.ReadUInt16()

	// Read the actual message string
	message := data.ReadString(uint(messageLength))

	fmt.Println(message)
	// Prepare the response buffer to broadcast the message
	bf := serializers.NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_MESSAGE_SAY) // Server opcode for broadcasting the message
	bf.WriteUInt32(player.ID)                              // Player's ID who sent the message
	bf.WriteUInt16(messageLength)                          // Include the message length
	bf.WriteString(message)                                // Write the actual message string

	// Broadcast the message to all players in the lobby
	player.BroadcastLobby(bf)
}
