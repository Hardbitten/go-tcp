package deserializers

import (
	"main/utils"
)

func DeserializePlayerCastSpell(data *utils.ByteBuffer) (uint32, utils.Vector3, error) {
	SpellID := data.ReadUInt32()
	position := utils.NewVector(data.ReadFloat(), data.ReadFloat(), data.ReadFloat())
	return SpellID, position, nil
}
