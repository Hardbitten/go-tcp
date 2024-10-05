package deserializers

import (
	"bytes"
	"main/utils"
)

// DeserializeAuthLogin deserializes data for AUTH_LOGIN opcode
func DeserializeAuthLogin(data *utils.ByteBuffer) (string, string, error) {
	// Read username
	usernameChunk := make([]byte, 30)
	data.GetCurrentStream().Read(usernameChunk)
	username := string(bytes.TrimSpace(usernameChunk))

	passwordChunk := make([]byte, 30)
	data.GetCurrentStream().Read(passwordChunk)
	password := string(bytes.TrimSpace(passwordChunk))

	return username, password, nil

}
