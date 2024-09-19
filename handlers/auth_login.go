package handlers

import (
	"fmt"
	"main/models"
	"main/serializers"
	// Import other necessary packages
)

func HandleAuthLogin(data *serializers.ByteBuffer, player *models.Player) {
	// Deserialize data
	// Process login
	fmt.Println("Handling AUTH_LOGIN")
	username, password := serializers.DeserializeAuthLogin(data)
	fmt.Printf("Handling AUTH_LOGIN for user: %s with password: %s\n", username, password)

	fmt.Println(username)
	fmt.Println(password)

	// Check user and pass in db.
}
