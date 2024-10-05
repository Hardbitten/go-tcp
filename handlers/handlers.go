package handlers

import (
	"main/models"
	op "main/opcodes"
	"main/utils"
)

// Define a type for the handler function
type HandlerType func(data *utils.ByteBuffer, player *models.Player)

var Handlers = map[uint16]HandlerType{
	op.CMSG_OPCODE_AUTH_LOGIN:             HandleAuthLogin,
	op.CMSG_OPCODE_MATCH_FIND_START:       HandleMatchFindStart,
	op.CMSG_OPCODE_MATCH_WARM_UP:          HandleMatchWarmUp,
	op.CMSG_OPCODE_MATCH_FIND_CANCEL:      HandleMatchFindCancel,
	op.CMSG_OPCODE_PLAYER_ENTER_WORLD:     HandlePlayerEnterWorld,
	op.CMSG_OPCODE_PLAYER_DISCONNECT:      HandlePlayerDisconnect,
	op.CMSG_OPCODE_PLAYER_UPDATE_POSITION: HandlePlayerUpdatePosition,
	op.CMSG_OPCODE_PLAYER_UPDATE_ROTATION: HandlePlayerUpdateRotation,
	op.CMSG_OPCODE_MESSAGE_CHAT:           HandlePlayerMessageChat,
}
