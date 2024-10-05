package serializers

import (
	"main/opcodes"
	utils "main/utils"
)

func SerializePlayerPosition(PlayerID uint32, Position utils.Vector3) *ByteBuffer {

	bf := NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_UPDATE_POSITION)

	bf.WriteUInt32(PlayerID)

	X := Position.X
	Y := Position.Y
	Z := Position.Z

	bf.WriteFloat(X) // p X
	bf.WriteFloat(Y) // p Y
	bf.WriteFloat(Z) // p Z

	return bf
}
