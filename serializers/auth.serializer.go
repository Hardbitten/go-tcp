package serializers

import "bytes"

// DeserializeAuthLogin deserializes data for AUTH_LOGIN opcode
func DeserializeAuthLogin(data *ByteBuffer) (string, string) {
	// Read username
	username := make([]byte, 30)
	data.readStream.Read(username)

	// Read password
	password := make([]byte, 30)
	data.readStream.Read(password)

	return string(bytes.Trim(username, "\x00")), string(bytes.Trim(password, "\x00"))

}
