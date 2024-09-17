package serializers

import (
	"main/models"
)

// SerializePlayer serializes a Player model into a ByteBuffer
func SerializePlayer(buffer *ByteBuffer, player models.Player) {
	buffer.WriteUInt32(player.ID)
	SerializePosition(buffer, player.Position)
}

// DeserializePlayer deserializes a Player model from a ByteBuffer
func DeserializePlayer(buffer *ByteBuffer) models.Player {
	player := models.Player{
		ID:       buffer.ReadUInt32(), // Assuming a fixed length for player ID
		Position: DeserializePosition(buffer),
	}
	return player
}
