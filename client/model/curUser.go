package model

import (
	"exercise/chatroom/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	User message.User
}
