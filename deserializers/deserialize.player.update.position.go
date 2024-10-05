package deserializers

import "main/utils"

func DeserializePlayerUpdatePosition(data *utils.ByteBuffer) (utils.Vector3, error) {
	position := utils.NewVector(data.ReadFloat(), data.ReadFloat(), data.ReadFloat())
	return position, nil
}
