package main

import (
	"exercise/chatroom/client/process"
	"fmt"
)

var (
	userId   int
	userPwd  string
	userName string
)

func main() {
	//接收用户选择
	k := 0

	for {
		fmt.Println("----------欢迎登陆多人聊天系统----------")
		fmt.Println("\t\t 1 登录聊天室")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出系统")
		fmt.Println("\t\t 请选择(1~3)")

		fmt.Scanf("%d\n", &k)

		switch k {
		case 1:
			fmt.Println("----------登录聊天室----------")
			fmt.Println("请输入用户Id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)

			up := process.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				fmt.Println("login err=", err)
			}
		case 2:
			fmt.Println("----------注册用户----------")
			fmt.Println("请输入用户Id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名称:")
			fmt.Scanf("%s\n", &userName)

			up := process.UserProcess{}
			err := up.Register(userId, userPwd, userName)
			if err != nil {
				fmt.Println("Register err=", err)
			}
		case 3:
			fmt.Println("退出系统")
			return
		default:
			fmt.Println("输入有误,请重新输入...")
		}
	}

}
