package serializers

import "main/models"

// SerializePlayerEnterWorld serializes data for PLAYER_ENTER_WORLD opcode
func SerializePlayerEnterWorld(player *models.Player) *ByteBuffer {
	buffer := NewByteBuffer()

	// Write opcode for PLAYER_ENTER_WORLD (assume it's 0x01)
	buffer.WriteUInt16(0x01) // Replace with the correct opcode

	// Serialize player ID (4 bytes for int32)
	buffer.WriteUInt32(player.ID)

	// Serialize player position (3 * 4 bytes for float32 each)
	buffer.WriteFloat(player.Position.X)
	buffer.WriteFloat(player.Position.Y)
	buffer.WriteFloat(player.Position.Z)

	// Serialize other player data as needed (e.g., name, status)
	// For example, serialize player name with fixed length
	name := player.Name
	if len(name) > 30 {
		name = name[:30] // truncate if necessary
	}
	buffer.WriteBytesWithLength([]byte(name), 30) // Fixed length of 30 bytes

	return buffer
}
