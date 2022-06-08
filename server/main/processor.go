package main

import (
	"exercise/chatroom/common/message"
	"exercise/chatroom/server/processes"
	"exercise/chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (p *Processor) serverProcessMes(mes *message.Message) (err error) {

	fmt.Println("mes====", mes)
	switch mes.Type {
	case message.LoginMesType:
		up := &processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerLoginMes(mes)
		if err != nil {
			fmt.Println("serverLoginMes err", err)
		}
	case message.RegisterMesType:
		up := &processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerRegisterMes(mes)
		if err != nil {
			fmt.Println("serverRegisterMes err", err)
		}
	case message.SmsMesType:
		sp := &processes.SmsProcess{}
		sp.SendGroupMes(mes)
	case message.PerMesType:
		sp := &processes.SmsProcess{}
		sp.SendPerMes(mes)
	default:
		fmt.Println("不存在的消息类型,无法处理...")
	}
	return
}

func (p *Processor) process2() (err error) {
	for {
		tf := utils.Transfer{
			Conn: p.Conn,
		}
		fmt.Println("读取客户端发送的数据...")
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出....")
				return err
			} else {
				fmt.Println("ReadPkg err=", err)
				return err
			}

		}
		err = p.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("erverProcessMes err", err)
			return err
		}
	}
}
