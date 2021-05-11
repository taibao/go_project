package main

import (
	"chat/client"
	"chat/common/message"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

//处理登录请求
func serverProcessLogin(conn net.Conn,mes *message.Message)(err error){
	//从mes中取出mes.Data
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil{
		fmt.Println("json.Unmarshal fail err=",err)
		return
	}

	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//再声明一个LoginResMes,并完成赋值
	var loginResMes message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPwd == "123456"{
		//合法
		loginResMes.Code = 200
	}else{
		//不合法
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用..."
	}

	//将loginResMes序列化
	data,err := json.Marshal(loginResMes)
	if err != nil{
		fmt.Println("loginResMes序列化失败",err)
		return
	}
	//将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes 进行序列化，准备发送
	data,err = json.Marshal(resMes)
	if err != nil{
		fmt.Println("loginResMes序列化失败",err)
		return
	}
	//6: 发送data,我们将其封装到writePkg函数
	err = client.WritePkg(conn,data)
	return
}


//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn,mes *message.Message) (err error){

	switch mes.Type{
	case message.LoginMesType:
		//处理登录逻辑
		err = serverProcessLogin(conn,mes)
	case message.RegisterMesType:
		//处理注册逻辑
	default:
		fmt.Println("消息类型不存在，无法处理。。。")
	}
	return
}


//处理和客户端的通讯
func process(conn net.Conn){
	//需要延时关闭conn
	defer conn.Close()
	//读客户端发送的信息
	for{
		//这里我们将读取数据包，直接封装成一个函数readPkg，返回Message，Err
		mes,err := client.ReadPkg(conn)
		if err != nil{
			if err ==io.EOF{
				fmt.Println("客户端退出，服务器端也退出")
				return
			}else{
				fmt.Println("readpkg err=",err)
				return
			}
		}
		//fmt.Println("消息内容 mes=",mes)
		err = serverProcessMes(conn,&mes)
		if err != nil{
			return
		}
	}
}

func main(){
	//提示信息
	fmt.Println("服务器在8889端口监听。。。")
	listen,err := net.Listen("tcp","0.0.0.0:8889") //广播8889端口
	defer listen.Close()

	if err != nil{
		fmt.Println("net.listen err=",err)
		return
	}

	//一旦监听成功，就等待客户端来连接服务器端
	for{
		fmt.Println("等待客户端连接")
		conn,err := listen.Accept()
		if err !=nil{
			fmt.Println("listen.Accept err=",err)
		}

		//一旦连接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}


}