package enums

type AuthHandshakeStatus uint16

// Define authentication handshake results
const (
	AUTH_HANDSHAKE_RESULT_OK        AuthHandshakeStatus = 0x00
	AUTH_HANDSHAKE_RESULT_ERR       AuthHandshakeStatus = 0x01
	AUTH_HANDSHAKE_RESULT_NOT_FOUND AuthHandshakeStatus = 0x02
	AUTH_HANDSHAKE_RESULT_BANNED    AuthHandshakeStatus = 0x03
)
