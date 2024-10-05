package deserializers

import "main/utils"

func DeserializePlayerUpdateRotation(data *utils.ByteBuffer) (float32, error) {
	rotation := data.ReadFloat()
	return rotation, nil
}
