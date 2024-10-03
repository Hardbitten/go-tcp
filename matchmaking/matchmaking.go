package matchmaking

import (
	"fmt"
	"main/models"
	"sync"
)

var (
	nextLobbyID uint32     = 1 // Start from 1 for normal lobbies
	mu          sync.Mutex     // Mutex for safe concurrent access
)

type LobbyManager struct {
	lobbies map[uint32]*models.Lobby // Use uint32 as the key for lobbies
}

var lobbyManager = &LobbyManager{
	lobbies: make(map[uint32]*models.Lobby),
}

// FindOrCreateLobby finds an existing lobby or creates a new one
func FindOrCreateLobby(player *models.Player) *models.Lobby {
	mu.Lock()
	defer mu.Unlock()

	// Find an existing lobby with fewer than 4 players
	for id, lobby := range lobbyManager.lobbies {
		if len(lobby.Players) < 4 && id != 0 { // Avoid warm-up lobby (ID = 0)
			fmt.Printf("Joining existing lobby %d\n", id)
			return lobby
		}
	}

	// No existing lobby found, create a new one
	newLobbyID := nextLobbyID
	nextLobbyID++
	newLobby := models.NewLobby([]*models.Player{})
	lobbyManager.lobbies[newLobbyID] = newLobby
	fmt.Println("Created new lobby:", newLobbyID)
	return newLobby
}

// AddPlayerToMatchmaking adds a player to a lobby, either existing or newly created
func AddPlayerToMatchmaking(player *models.Player, isWarmup bool) {
	var lobby *models.Lobby

	if isWarmup {
		// Always join the warm-up lobby with ID = 0
		lobby = getOrCreateWarmupLobby()
	} else {
		// Find or create a normal lobby
		lobby = FindOrCreateLobby(player)
	}

	// Add the player to the found or created lobby
	lobby.AddPlayer(player)
	fmt.Printf("Player added to lobby %d: %s\n", lobby.ID, player.Name)

	player.Lobby = lobby
	// Warm-up lobby behavior
	if lobby.ID == 0 {
		// Send MatchReady immediately for warm-up lobbies
		bf := models.SerializeMatchReady(player)
		player.Session.Conn.Write(bf.GetData())
	} else {
		// For normal lobbies, wait for enough players (e.g., 4 players) to start the match
		if len(lobby.Players) >= 4 {
			startMatch(lobby)
		}
	}
}

// getOrCreateWarmupLobby retrieves or creates a warm-up lobby with ID 0
func getOrCreateWarmupLobby() *models.Lobby {
	mu.Lock()
	defer mu.Unlock()

	// Check if warm-up lobby already exists
	warmupLobby, exists := lobbyManager.lobbies[0]
	if !exists {
		// Create the warm-up lobby if it doesn't exist
		warmupLobby = models.NewLobby([]*models.Player{})
		lobbyManager.lobbies[0] = warmupLobby
		fmt.Println("Created warm-up lobby (ID = 0)")
	}
	return warmupLobby
}

// startMatch begins the game once the lobby has enough players
func startMatch(lobby *models.Lobby) {
	fmt.Printf("Starting match for lobby ID: %d with %d players\n", lobby.ID, len(lobby.Players))

	// Send ready check to all players
	for _, player := range lobby.Players {
		// Set the lobby for each player
		player.Lobby = lobby

		// Perform the ready check for each player
		player.Session.SendReadyCheck()
	}

	// Simulate logic for checking if all players are ready
	// You should replace this with actual ready check logic in your game
	allPlayersReady := true // Placeholder: implement your ready check logic
	if allPlayersReady {
		// If all players are ready, start the game
		lobby.StartGame()
		fmt.Printf("Match started for lobby ID: %d\n", lobby.ID)

		// Notify all players that the match has started
		for _, player := range lobby.Players {
			bf := models.SerializeMatchReady(player)
			player.Session.Conn.Write(bf.GetData())
		}
	} else {
		fmt.Printf("Not all players are ready in lobby ID: %d\n", lobby.ID)
	}
}
