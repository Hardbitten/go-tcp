package models

import (
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

// NewPlayer creates a new player and associates it with a session
func NewPlayer(session *Session) *Player {
	LastPlayerId++
	return &Player{
		ID:       LastPlayerId,
		Session:  session,
		Position: Vector3{X: 0, Y: 0, Z: 0}, // Initialize position at origin
		Rotation: 0,
	}
}

// SerializePlayerEnterWorld serializes data for PLAYER_ENTER_WORLD opcode
func SerializePlayerEnterWorld(player *Player) *serializers.ByteBuffer {
	buffer := serializers.NewByteBuffer()

	// Write opcode for PLAYER_ENTER_WORLD (assume it's 0x01)
	buffer.WriteUInt16(opcodes.SMSG_OPCODE_PLAYER_ENTER_WORLD) // Replace with the correct opcode

	// Serialize player ID (4 bytes)
	buffer.WriteUInt32(player.ID)

	// Serialize player position (3 * 4 bytes for float32 each)
	buffer.WriteFloat(player.Position.X)
	buffer.WriteFloat(player.Position.Y)
	buffer.WriteFloat(player.Position.Z)

	buffer.WriteFloat(player.Rotation)

	// Serialize other player data as needed (e.g., name, status)
	// For example, serialize player name with fixed length
	name := player.Name
	if len(name) > 30 {
		name = name[:30] // truncate if necessary
	}
	buffer.WriteBytesWithLength([]byte(name), 30) // Fixed length of 30 bytes

	return buffer
}

func (player *Player) BroadcastLobby(bf *serializers.ByteBuffer) {
	for _, lobyPlayer := range player.Lobby.Players {
		if lobyPlayer.ID != player.ID { // Don't send to the sender
			lobyPlayer.Session.Conn.Write(bf.GetData())
		}
	}
}
