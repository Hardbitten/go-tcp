package main

// Define opcodes
const (
	// Auth opcodes
	OPCODE_AUTH_LOGIN    uint16 = 0x01
	OPCODE_AUTH_REGISTER uint16 = 0x02

	// Game opcodes
	OPCODE_GAME_ENTER_WORLD     uint16 = 0x100
	OPCODE_GAME_UPDATE_POSITION uint16 = 0x101
	OPCODE_GAME_MESSAGE_SAY     uint16 = 0x102
)
