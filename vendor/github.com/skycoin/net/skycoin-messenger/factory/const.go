package factory

const (
	MSG_OP_SIZE         = 1
	MSG_PUBLIC_KEY_SIZE = 33
)

const (
	MSG_HEADER_BEGIN = 0
	MSG_OP_BEGIN
	MSG_OP_END = MSG_OP_BEGIN + MSG_OP_SIZE
	MSG_HEADER_END
)

const (
	SEND_MSG_META_BEGIN = MSG_HEADER_END

	SEND_MSG_PUBLIC_KEY_BEGIN
	SEND_MSG_PUBLIC_KEY_END = SEND_MSG_PUBLIC_KEY_BEGIN + MSG_PUBLIC_KEY_SIZE

	SEND_MSG_TO_PUBLIC_KEY_BEGIN
	SEND_MSG_TO_PUBLIC_KEY_END = SEND_MSG_TO_PUBLIC_KEY_BEGIN + MSG_PUBLIC_KEY_SIZE

	SEND_MSG_META_END
)

const (
	OP_REG = iota
	OP_SEND
	OP_CUSTOM
	OP_OFFER_SERVICE
	OP_QUERY_SERVICE_NODES
	OP_QUERY_BY_ATTRS
	OP_BUILD_APP_CONN
	OP_FORWARD_NODE_CONN
	OP_BUILD_NODE_CONN
	OP_FORWARD_NODE_CONN_RESP
	OP_BUILD_APP_CONN_OK

	OP_SIZE
)

const RESP_PREFIX = 0x80
