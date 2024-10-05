package serializers

import (
	"main/opcodes"
	"main/utils"
)

func SerializePlayerSync(PlayerID uint32, Position utils.Vector3, Rotation float32, Name string) *utils.ByteBuffer {
	// Get latest Data for user
	bfSync := utils.NewByteBuffer()
	bfSync.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_SYNC)
	bfSync.WriteUInt32(PlayerID)

	bfSync.WriteFloat(Position.X)
	bfSync.WriteFloat(Position.Y)
	bfSync.WriteFloat(Position.Z)

	bfSync.WriteFloat(Rotation)

	name := Name
	if len(name) > 30 {
		name = name[:30] // truncate
	}
	bfSync.WriteBytesWithLength([]byte(name), 30) // Fixed length of 30 bytes
	return bfSync
}
