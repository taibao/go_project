package main

import (
	"fmt"
)

//定义两个变量，一个表示用户id 一个表示用户密码
var userId int
var userPwd string

func main() {

	var key int
	var loop= true

	for loop{
		fmt.Println("-------欢迎登录多人聊天系统-----------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 4 请选择（1-3）：")
		fmt.Scanf("%d\n",&key)
		switch key{
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("退出系统")
		default:
			fmt.Println("你的输入有错")
		}
	}

	//更新用户的输入，显示新的提示信息
	if key == 1{
		//说明用户要登录
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d\n",&userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%d\n",&userPwd)
		err := login(userId,userPwd)
		if err != nil{
			fmt.Println("登陆失败")
		}else{
			fmt.Println("登录成功")
		}
	}else if key == 2{
		fmt.Println("进行用户注册")
	}


}
