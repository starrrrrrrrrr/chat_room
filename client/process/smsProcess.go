package process

import (
	"encoding/json"
	"exercise/chatroom/client/utils"
	"exercise/chatroom/common/message"
	"fmt"
)

type SmsProcess struct {
}

func (s *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.User.UserId
	smsMes.UserStatus = CurUser.User.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("marshal err=", err)
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("write err=", err)
		return
	}

	return
}

func (s *SmsProcess) SendPerMes(content string, Id int) (err error) {
	var mes message.Message
	mes.Type = message.PerMesType

	var perMes message.PerMes
	perMes.Content = content
	perMes.UserId = CurUser.User.UserId
	perMes.UserStatus = CurUser.User.UserStatus
	perMes.Id = Id

	data, err := json.Marshal(perMes)
	if err != nil {
		fmt.Println("marshal err=", err)
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("write err=", err)
		return
	}

	return
}
