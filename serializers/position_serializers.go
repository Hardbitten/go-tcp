package serializers

import (
	"main/models"
)

// SerializePosition serializes a Position model into a ByteBuffer
func SerializePosition(buffer *ByteBuffer, position models.Position) {
	buffer.WriteFloat(position.X)
	buffer.WriteFloat(position.Y)
	buffer.WriteFloat(position.Z)
	buffer.WriteFloat(position.O)
}

// DeserializePosition deserializes a Position model from a ByteBuffer
func DeserializePosition(buffer *ByteBuffer) models.Position {
	position := models.Position{
		X: buffer.ReadFloat(),
		Y: buffer.ReadFloat(),
		Z: buffer.ReadFloat(),
		O: buffer.ReadFloat(),
	}
	return position
}
