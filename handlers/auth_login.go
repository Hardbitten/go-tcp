package handlers

import (
	"fmt"
	"main/deserializers"
	"main/enums"
	"main/models"
	"main/serializers"
	"main/utils"
)

func HandleAuthLogin(data *utils.ByteBuffer, player *models.Player) {
	fmt.Println("Handling AUTH_LOGIN")
	username, password, err := deserializers.DeserializeAuthLogin(data)
	if err != nil {
		fmt.Printf("Error ! [%s]", err)
	}

	fmt.Printf("Handling AUTH_LOGIN for user: %s with password: %s\n", username, password)

	// Check user and pass in db.
	bf := serializers.SerializeAuthHandshake(data, enums.AUTH_HANDSHAKE_RESULT_OK)
	player.Session.Conn.Write(bf.GetData())

}
