package process

import (
	"encoding/json"
	"exercise/chatroom/common/message"
	"fmt"
)

func OutSmsMes(mes *message.Message) {
	var sms message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &sms)
	if err != nil {
		fmt.Println("outSmsMes err=", err)
		return
	}

	info := fmt.Sprintf("用户(Id:%d):\n%s", sms.UserId, sms.Content)
	fmt.Println(info)
}

func OutPerMes(mes *message.Message) {
	var per message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &per)
	if err != nil {
		fmt.Println("outPerMes err=", err)
		return
	}

	info := fmt.Sprintf("用户(Id:%d)对你说:\n%s", per.UserId, per.Content)
	fmt.Println(info)
}
