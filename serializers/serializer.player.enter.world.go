package serializers

import (
	"main/opcodes"
	"main/utils"
)

func SerializePlayerEnterWorld(PlayerID uint32, Position utils.Vector3, Rotation float32, Name string) *utils.ByteBuffer {
	buffer := utils.NewByteBuffer()

	buffer.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_ENTER_WORLD)

	buffer.WriteUInt32(PlayerID)

	buffer.WriteFloat(Position.X)
	buffer.WriteFloat(Position.Y)
	buffer.WriteFloat(Position.Z)

	buffer.WriteFloat(Rotation)

	name := Name
	if len(name) > 30 {
		name = name[:30] // truncate
	}
	buffer.WriteBytesWithLength([]byte(name), 30) // Fixed length of 30 bytes

	return buffer
}
