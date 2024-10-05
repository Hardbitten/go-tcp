package models

import (
	"fmt"
	op "main/opcodes"
	"main/utils"
)

type Lobby struct {
	ID      uint32
	Players []*Player
}

func NewLobby(players []*Player) *Lobby {
	return &Lobby{
		Players: players,
	}
}

func (l *Lobby) AddPlayer(player *Player) {
	l.Players = append(l.Players, player)
}

func (l *Lobby) RemovePlayer(playerID uint32) {
	if len(l.Players) == 0 {
		fmt.Println("No players in the lobby to remove.")
		return
	}

	for i, p := range l.Players {
		if p == nil {
			continue // Skip nil players
		}
		if p.ID == playerID {
			// Remove the player from the slice
			l.Players = append(l.Players[:i], l.Players[i+1:]...)
			fmt.Printf("Player %d removed from the lobby.\n", playerID)
			return
		}
	}

	fmt.Printf("Player %d not found in the lobby.\n", playerID)
}

func (l *Lobby) StartGame() {
	fmt.Println("Starting game for lobby with players:", len(l.Players))

	for _, player := range l.Players {
		bf := SerializeMatchReady(player)

		player.Session.Conn.Write(bf.GetData())
	}
}

func SerializeMatchReady(player *Player) *utils.ByteBuffer {
	buffer := utils.NewByteBuffer()
	buffer.WriteUInt16(op.SMSG_OPCODE_MATCH_READY)
	buffer.WriteUInt32(player.ID)

	return buffer
}
