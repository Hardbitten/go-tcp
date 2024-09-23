package models

import (
	"fmt"
	"net"
	"sync"
)

var (
	lastSessionID = 0
	sessionIDLock sync.Mutex // To make ID generation thread-safe
)

// Session represents a user's connection
type Session struct {
	ID   string
	Conn net.Conn
}

// NewSession creates a new session for a player
func NewSession(conn net.Conn) *Session {
	sessionIDLock.Lock()
	defer sessionIDLock.Unlock()

	lastSessionID++
	return &Session{
		ID:   generateUniqueID(),
		Conn: conn,
	}
}

// SendReadyCheck sends a ready check message to the player
func (s *Session) SendReadyCheck() {
	fmt.Println("Sending ready check to player:", s.ID)
	// Implement message sending logic here
}

func generateUniqueID() string {
	return fmt.Sprintf("player_%d", lastSessionID)
}
