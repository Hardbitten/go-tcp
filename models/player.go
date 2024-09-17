package models

type Player struct {
	Session  *Session // Player inherits the session to access socket and ID
	Position Vector3
	Lobby    *Lobby // Player is part of a lobby when matchmaking is done
}

// NewPlayer creates a new player and associates it with a session
func NewPlayer(session *Session) *Player {
	return &Player{
		Session:  session,
		Position: Vector3{X: 0, Y: 0, Z: 0}, // Initialize position at origin
	}
}
