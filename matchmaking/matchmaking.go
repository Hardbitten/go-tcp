package matchmaking

import (
	"fmt"
	"main/models"
)

var waitingPlayers []*models.Player

// AddPlayerToMatchmaking adds a player to the matchmaking queue
func AddPlayerToMatchmaking(player *models.Player) {
	waitingPlayers = append(waitingPlayers, player)
	if len(waitingPlayers) >= 4 { // Example: Start match with 4 players
		startMatch()
	}
}

// startMatch starts a match with the waiting players
func startMatch() {
	fmt.Println("Starting match with players:", len(waitingPlayers))

	// Send ready check to all players
	for _, session := range waitingPlayers {
		session.Session.SendReadyCheck()
	}

	// If all players are ready, move them to a lobby
	// Example: after some ready check confirmation logic
	// Create a new lobby and assign players
	lobby := models.NewLobby(waitingPlayers)
	lobby.StartGame()

	// Clear matchmaking queue
	waitingPlayers = []*models.Player{}
}
