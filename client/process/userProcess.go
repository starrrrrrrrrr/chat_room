package process

import (
	"encoding/json"
	"exercise/chatroom/client/utils"
	"exercise/chatroom/common/message"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
}

func (up *UserProcess) Login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", hostIP)
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.LoginMesType
	var loginmes message.LoginMes
	loginmes.UserId = userId
	loginmes.UserPwd = userPwd

	data, err := json.Marshal(loginmes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("son.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("ReadPkg err", err)
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
	}

	if loginResMes.Code == 200 {
		fmt.Println("登录成功...")

		CurUser.Conn = conn
		CurUser.User.UserId = userId
		CurUser.User.UserStatus = message.UserOnline

		fmt.Println("当前在线用户")
		for _, v := range loginResMes.Users {
			fmt.Println("用户id:", v)

			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}

			OnlineUsers[v] = user
		}

		go serverProcessMes(conn)

		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}

func (up *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", hostIP)
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType
	var register message.RegisterMes
	register.User.UserId = userId
	register.User.UserPwd = userPwd
	register.User.UserName = userName

	data, err := json.Marshal(register)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("son.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("ReadPkg err", err)
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
	}

	if registerResMes.Code == 200 {
		fmt.Println("注册成功,可以前往登录...")
		return
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}

	return
}
