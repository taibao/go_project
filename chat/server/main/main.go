package main

import (
	"chat/server/model"
	"chat/server/processes"
	"chat/server/utils"
	"fmt"
	"net"
	"time"
)

//处理和客户端的通讯
func process(conn net.Conn){
	//需要延时关闭conn
	defer conn.Close()
	//调用主控
	processor := &processes.Processor{
		Conn : conn,
	}
	err := processor.Process2()
	if err != nil{
		fmt.Println("客户端和服务器通讯协程错误=err",err)
		return
	}

}

//编写函数完成对UserDao的初始化任务
func initUserDao(){
	//这里的pool本身就是一个全局的变量
	//需要注意一个初始化顺序的问题
	//initPool
	model.MyUserDao = model.NewUserDao(utils.Pool)
}

func main(){
	//当服务器启动时
	utils.InitPool("localhost:6379",16,0,300*time.Second)
	initUserDao()

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