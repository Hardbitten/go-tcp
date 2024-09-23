package handlers

import (
	"bytes"
	"fmt"
	"main/models"
	"main/opcodes"
	"main/serializers"
	// Import other necessary packages
)

func HandleAuthLogin(data *serializers.ByteBuffer, player *models.Player) {
	// Deserialize data
	// Process login
	fmt.Println("Handling AUTH_LOGIN")
	username, password := DeserializeAuthLogin(data)
	fmt.Printf("Handling AUTH_LOGIN for user: %s with password: %s\n", username, password)

	fmt.Println(username)
	fmt.Println(password)

	// Check user and pass in db.

	bf := SerializeAuthHandshake()
	player.Session.Conn.Write(bf.GetData())

}

// DeserializeAuthLogin deserializes data for AUTH_LOGIN opcode
func DeserializeAuthLogin(data *serializers.ByteBuffer) (string, string) {
	// Read username
	username := make([]byte, 30)
	data.GetCurrentStream().Read(username)

	// Read password
	password := make([]byte, 30)
	data.GetCurrentStream().Read(password)

	return string(bytes.Trim(username, "\x00")), string(bytes.Trim(password, "\x00"))

}

// DeserializeAuthLogin deserializes data for AUTH_LOGIN opcode
func SerializeAuthHandshake() (data *serializers.ByteBuffer) {
	bf := serializers.NewByteBuffer()
	bf.WriteUInt16(opcodes.SMSG_OPCODE_AUTH_HANDSHAKE)

	bf.WriteUInt16(opcodes.AUTH_HANDSHAKE_RESULT_OK)

	return bf
}
