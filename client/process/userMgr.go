package process

import (
	"exercise/chatroom/client/model"
	"exercise/chatroom/common/message"
	"fmt"
)

var OnlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

const hostIP = "127.0.0.1:9999"

func OutPutOnlineUsers() {
	fmt.Println("----------当前在线用户列表----------")
	for id := range OnlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

func UpdateUserStatus(notify *message.NotifyUserStatusMes) {
	user, ok := OnlineUsers[notify.UserId]
	if !ok {
		user = &message.User{
			UserId: notify.UserId,
		}
	}

	user.UserStatus = notify.UserStatus
	OnlineUsers[notify.UserId] = user

	OutPutOnlineUsers()
}
