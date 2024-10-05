package serializers

import (
	"main/opcodes"
	"main/utils"
)

func SerializePlayerCastSpell(PlayerID uint32, SpellID uint32, TargetPosition utils.Vector3) *utils.ByteBuffer {
	bf := utils.NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_UPDATE_POSITION)
	bf.WriteUInt32(PlayerID)
	bf.WriteUInt32(SpellID)
	bf.WriteFloat(TargetPosition.X)
	bf.WriteFloat(TargetPosition.Y)
	bf.WriteFloat(TargetPosition.Z)
	return bf
}
