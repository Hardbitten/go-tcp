package models

import (
	"fmt"
	op "main/opcodes"
	"main/serializers"
)

type Lobby struct {
	Players []*Player
}

// NewLobby creates a new lobby with players
func NewLobby(players []*Player) *Lobby {
	return &Lobby{
		Players: players,
	}
}

// StartGame starts the game for all players in the lobby
func (l *Lobby) StartGame() {
	fmt.Println("Starting game for lobby with players:", len(l.Players))
	// Implement game start logic, broadcast data to all players

	for _, player := range l.Players {
		bf := SerializeMatchReady(player)

		player.Session.Conn.Write(bf.GetData())
	}
}

func SerializeMatchReady(player *Player) *serializers.ByteBuffer {
	buffer := serializers.NewByteBuffer()
	buffer.WriteUInt16(op.SMSG_OPCODE_MATCH_READY)

	return buffer
}
