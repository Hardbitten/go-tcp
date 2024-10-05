package models

import (
	"fmt"
	"main/utils"
)

type Player struct {
	ID       uint32
	Name     string
	Session  *Session // Player inherits the session to access socket and ID
	Position utils.Vector3
	Rotation float32
	Lobby    *Lobby // Player is part of a lobby when matchmaking is done

	Health  int
	Mana    int
	ClassID ClassType
}

var LastPlayerId uint32 = 0

func NewPlayer(session *Session) *Player {
	LastPlayerId++
	return &Player{
		ID:       LastPlayerId,
		Session:  session,
		Position: utils.Vector3{X: 0, Y: 0, Z: 0}, // Initialize position at origin
		Rotation: 0,
	}
}

// The Cast method now takes a Player as the caster
func (caster *Player) CastSpell(s *Spell) bool {
	if caster.Mana < s.ManaCost {
		fmt.Printf("%s does not have enough mana to cast %s\n", caster.Name, s.Name)
		return false
	}

	// Apply spell effects, such as reducing mana and dealing damage to targets
	fmt.Printf("%s casts %s!\n", caster.Name, s.Name)

	// Deduct mana from the player
	caster.Mana -= s.ManaCost

	// Handle additional logic here like cooldowns, applying effects to enemies, etc.

	return true
}

func (player *Player) BroadcastLobby(bf *utils.ByteBuffer) {
	// Check if the player is in a valid lobby
	if player.Lobby == nil {
		fmt.Println("Error: Player is not assigned to any lobby.")
		return
	}

	// Check if the lobby has any players
	if len(player.Lobby.Players) == 0 {
		fmt.Println("Error: Lobby has no players.")
		return
	}

	// Iterate through the lobby players and broadcast the message
	for _, lobyPlayer := range player.Lobby.Players {
		// Skip sending the message to the sender (the player who initiated the broadcast)
		if lobyPlayer.ID != player.ID {
			// Check if the player's session or connection is nil
			if lobyPlayer.Session == nil || lobyPlayer.Session.Conn == nil {
				fmt.Printf("Error: Session or connection for player %d is nil, skipping.\n", lobyPlayer.ID)
				continue
			}

			// Broadcast the message to the other players
			fmt.Printf("Broadcasting to player ID %d\n", lobyPlayer.ID)
			lobyPlayer.Session.Conn.Write(bf.GetData())
		}
	}
}
