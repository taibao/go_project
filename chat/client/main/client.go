package main

import (
	"chat/client/process"
	"fmt"
)


//定义两个变量，表示id和密码
var userId int
var userPwd string

func main(){
	//接收用户的选择
	var key int

	for true{
		fmt.Println("------------欢迎登录多人聊天系统--------------")
		fmt.Println("------------1：登录聊天室--------------")
		fmt.Println("------------2：注册用户--------------")
		fmt.Println("------------3：退出系统--------------")
		fmt.Println("------------4：请选择（1-3）--------------")

		fmt.Scanf("%d\n",&key)
		switch key {
		case 1:
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n",&userPwd)
			//完成登录
			//创建一个UserProcess的实例
			up := &process.UserProcess{}
			up.Login(userId,userPwd)
		case 2:
			fmt.Println("注册用户")

		case 3:
			fmt.Println("退出系统")
		default:
			fmt.Println("你的输入有误请重新输入")
		}
	}




}
