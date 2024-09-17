package models

import "fmt"

type Lobby struct {
	Players []*Session
}

// NewLobby creates a new lobby with players
func NewLobby(players []*Session) *Lobby {
	return &Lobby{
		Players: players,
	}
}

// StartGame starts the game for all players in the lobby
func (l *Lobby) StartGame() {
	fmt.Println("Starting game for lobby with players:", len(l.Players))
	// Implement game start logic, broadcast data to all players
}
