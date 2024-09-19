package main

// Define opcodes
const (
	// Auth opcodes
	OPCODE_AUTH_LOGIN    uint16 = 0x01
	OPCODE_AUTH_REGISTER uint16 = 0x02

	// Server Opcodes
	SMSG_OPCODE_PLAYER_ENTER_WORLD     uint16 = 0x100
	SMSG_OPCODE_PLAYER_UPDATE_POSITION uint16 = 0x101
	SMSG_OPCODE_PLAYER_MESSAGE_SAY     uint16 = 0x102

	// Client Opcodes
	CMSG_OPCODE_PLAYER_ENTER_WORLD uint16 = 0x500
)
