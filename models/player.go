package models

type Player struct {
	ID       uint32
	Name     string
	Session  *Session // Player inherits the session to access socket and ID
	Position Vector3
	Lobby    *Lobby // Player is part of a lobby when matchmaking is done
}

var LastPlayerId uint32 = 0

// NewPlayer creates a new player and associates it with a session
func NewPlayer(session *Session) *Player {
	LastPlayerId++
	return &Player{
		ID:       LastPlayerId,
		Session:  session,
		Position: Vector3{X: 0, Y: 0, Z: 0}, // Initialize position at origin
	}
}
