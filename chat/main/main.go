package main

import (
	"chat/client"
	"fmt"
)


//定义两个变量，表示id和密码
var userId int
var userPwd string

func main(){
	//接收用户的选择
	var key int
	//判断是否继续显示菜单
	var loop = true

	for loop{
		fmt.Println("------------欢迎登录多人聊天系统--------------")
		fmt.Println("------------1：登录聊天室--------------")
		fmt.Println("------------2：注册用户--------------")
		fmt.Println("------------3：退出系统--------------")
		fmt.Println("------------4：请选择（1-3）--------------")

		fmt.Scanf("%d\n",&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误请重新输入")
		}
	}

	//根据用户输入，显示新的提示信息
	if key == 1{
		//说明用户要登录
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d\n",&userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n",&userPwd)
		//先把登录的函数写到另一个文件
		err :=  client.Login(userId,userPwd)
		if err != nil{
			fmt.Println("登录失败")
		}else{
			fmt.Println("登录成功")
		}
	} else  if key ==2{
		fmt.Println("进行用户注册的逻辑...")
	}


}
