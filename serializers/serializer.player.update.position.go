package serializers

import (
	"main/opcodes"
	"main/utils"
)

func SerializePlayerPosition(PlayerID uint32, Position utils.Vector3) *utils.ByteBuffer {
	bf := utils.NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_UPDATE_POSITION)
	bf.WriteUInt32(PlayerID)
	bf.WriteFloat(Position.X) // p X
	bf.WriteFloat(Position.Y) // p Y
	bf.WriteFloat(Position.Z) // p Z
	return bf
}
