package models

import (
	"fmt"
	"main/opcodes"
	"main/serializers"
)

type Player struct {
	ID       uint32
	Name     string
	Session  *Session // Player inherits the session to access socket and ID
	Position Vector3
	Rotation float32
	Lobby    *Lobby // Player is part of a lobby when matchmaking is done
}

var LastPlayerId uint32 = 0

func NewPlayer(session *Session) *Player {
	LastPlayerId++
	return &Player{
		ID:       LastPlayerId,
		Session:  session,
		Position: Vector3{X: 0, Y: 0, Z: 0}, // Initialize position at origin
		Rotation: 0,
	}
}

func SerializePlayerEnterWorld(player *Player) *serializers.ByteBuffer {
	buffer := serializers.NewByteBuffer()

	buffer.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_ENTER_WORLD)

	buffer.WriteUInt32(player.ID)

	buffer.WriteFloat(player.Position.X)
	buffer.WriteFloat(player.Position.Y)
	buffer.WriteFloat(player.Position.Z)

	buffer.WriteFloat(player.Rotation)

	name := player.Name
	if len(name) > 30 {
		name = name[:30] // truncate
	}
	buffer.WriteBytesWithLength([]byte(name), 30) // Fixed length of 30 bytes

	return buffer
}

func (player *Player) BroadcastLobby(bf *serializers.ByteBuffer) {
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
