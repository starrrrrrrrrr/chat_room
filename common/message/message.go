package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
	PerMesType              = "PerMes"
)

const (
	UserOnline = iota
	UserOffline
	UserBusy
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId  int    `json:"userId"`
	UserPwd string `json:"userPwd"`
}

type LoginResMes struct {
	Code  int    `json:"code"` //500用户未注册，200登录成功
	Users []int  `json:"onlineUsers"`
	Error string `json:"error"`
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"` //500用户已经注册，200注册成功
	Error string `json:"error"`
}

type NotifyUserStatusMes struct {
	UserId     int `json:"userId"`
	UserStatus int `json:"userStatus"`
}

type SmsMes struct {
	User
	Content string
}

type PerMes struct {
	User
	Id      int
	Content string
}
