package main

import (
	"exercise/chatroom/server/model"
	"fmt"
	"net"
	"time"
)

// func ReadPkg(conn net.Conn) (mes message.Message, err error) {
// 	buf := make([]byte, 8096)
// 	fmt.Println("读取数据...")
// 	_, err = conn.Read(buf[0:4])
// 	if err != nil {
// 		fmt.Println("conn.Read err=", err)
// 		return
// 	}
// 	pkgLen := binary.BigEndian.Uint32(buf[0:4])
// 	n, err := conn.Read(buf[0:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Read err=", err)
// 		return
// 	}

// 	err = json.Unmarshal(buf[0:pkgLen], &mes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal err=", err)
// 		return
// 	}

// 	return
// }

// func WritePkg(conn net.Conn, data []byte) (err error) {
// 	pkgLen := uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
// 	n, err := conn.Write(buf[0:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write err=", err)
// 		return
// 	}

// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Write err=", err)
// 		return
// 	}

// 	return
// }

// func serverLoginMes(conn net.Conn, mes *message.Message) (err error) {
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		fmt.Println("serverLoginMes json.Unmarshal err=", err)
// 		return
// 	}

// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType
// 	var loginResMes message.LoginResMes

// 	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
// 		loginResMes.Code = 200
// 	} else {
// 		loginResMes.Code = 500
// 		loginResMes.Error = "用户未注册,请先注册..."
// 	}

// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal err=", err)
// 		return
// 	}
// 	resMes.Data = string(data)
// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal err=", err)
// 		return
// 	}

// 	err = utils.WritePkg(conn, data)
// 	if err != nil {
// 		fmt.Println("WritePkg err=", err)
// 	}
// 	return
// }

// func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 	case message.LoginMesType:
// 		err = serverLoginMes(conn, mes)
// 		if err != nil {
// 			fmt.Println("serverLoginMes err", err)
// 		}
// 	case message.RegisterMesType:

// 	default:
// 		fmt.Println("不存在的消息类型,无法处理...")
// 	}
// 	return
// }

func process(conn net.Conn) {
	defer conn.Close()

	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("协程错误err=", err)
		return
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	initPool(16, 0, 300*time.Second, "localhost:6379")
	initUserDao()

	fmt.Println("服务器在端口9999监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("服务器等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		go process(conn)
	}
}
