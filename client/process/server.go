package process

import (
	"encoding/json"
	"exercise/chatroom/client/utils"
	"exercise/chatroom/common/message"
	"fmt"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("---------- 恭喜登录成功 ----------")
	fmt.Println("----------1.显示在线用户----------")
	fmt.Println("----------2.发送群消息  ----------")
	fmt.Println("----------3.发送个人消息----------")
	fmt.Println("----------4.退出系统    ----------")
	fmt.Println("请选择(1-4):")
	var key int

	var content string
	smsProcess := &SmsProcess{}
	var Id int

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		OutPutOnlineUsers()
	case 2:
		fmt.Println("想对大家说什么:")
		fmt.Scanf("%s\n", &content)

		err := smsProcess.SendGroupMes(content)
		if err != nil {
			fmt.Println("SendGroupMes err=", err)
		}
	case 3:
		fmt.Println("请输入接收者的id:")
		fmt.Scanf("%d\n", &Id)
		fmt.Println("请输入想对ta说的话:")
		fmt.Scanf("%s\n", &content)

		err := smsProcess.SendPerMes(content, Id)
		if err != nil {
			fmt.Println("SendPerMes err=", err)
		}
	case 4:
		fmt.Println("...退出系统...")
		os.Exit(0)
	default:
		fmt.Println("选项不正确")
	}
}

func serverProcessMes(Conn net.Conn) {
	tf := utils.Transfer{
		Conn: Conn,
	}
	for {
		fmt.Println("读取服务端消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		fmt.Println("mes=", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notify message.NotifyUserStatusMes
			err := json.Unmarshal([]byte(mes.Data), &notify)
			if err != nil {
				fmt.Println("notify Unmarshal err=", err)
				return
			}

			UpdateUserStatus(&notify)
		case message.SmsMesType:
			OutSmsMes(&mes)
		case message.PerMesType:
			OutPerMes(&mes)
		default:
			fmt.Println("服务器发来了未知的消息...")
		}
	}
}
