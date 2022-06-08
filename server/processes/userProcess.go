package processes

import (
	"encoding/json"
	"exercise/chatroom/common/message"
	"exercise/chatroom/server/model"
	"exercise/chatroom/server/utils"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (u *UserProcess) ServerLoginMes(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("serverLoginMes json.Unmarshal err=", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器异常..."
		}
	} else {
		loginResMes.Code = 200
		u.UserId = loginMes.UserId
		userMgr.AddOnlineUser(u)
		u.NotifyOtherOnlineUsers(u.UserId)
		for id := range userMgr.OnlineUsers {
			loginResMes.Users = append(loginResMes.Users, id)
		}
		fmt.Println("登录成功,user=", user)
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg err=", err)
	}
	return
}

func (u *UserProcess) ServerRegisterMes(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("serverRegisterMes json.Unmarshal err=", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 500
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "服务器异常..."
		}
	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功,user=", registerMes.User)
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg err=", err)
	}
	return
}

func (u *UserProcess) NotifyOtherOnlineUsers(userId int) {

	for id, up := range userMgr.OnlineUsers {
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)
	}
}

func (u *UserProcess) NotifyMeOnline(userId int) {

	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notify message.NotifyUserStatusMes

	notify.UserId = userId
	notify.UserStatus = message.UserOnline

	data, err := json.Marshal(notify)
	if err != nil {
		fmt.Println("marshal notify err=", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal mes err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg err=", err)
		return
	}
}
