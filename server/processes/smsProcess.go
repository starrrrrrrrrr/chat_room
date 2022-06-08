package processes

import (
	"encoding/json"
	"exercise/chatroom/common/message"
	"exercise/chatroom/server/utils"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (s *SmsProcess) SendGroupMes(mes *message.Message) {
	var sms message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &sms)
	if err != nil {
		fmt.Println("sendGroupMes unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroupMes marshal err=", err)
		return
	}

	for id, up := range userMgr.OnlineUsers {
		if id == sms.UserId {
			continue
		}

		s.SendMesToEachOthers(data, up.Conn)
	}
}

func (s *SmsProcess) SendMesToEachOthers(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesToEachOthers writePkg err=", err)
	}
}

func (s *SmsProcess) SendPerMes(mes *message.Message) {
	var perMes message.PerMes
	err := json.Unmarshal([]byte(mes.Data), &perMes)
	if err != nil {
		fmt.Println("sendPerMes err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("sendGroupMes marshal err=", err)
		return
	}

	up := userMgr.OnlineUsers[perMes.Id]
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendPerMes writePkg err=", err)
	}
}
