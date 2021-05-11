package main

import (
	"chat/server/processes"
	"fmt"
	"net"
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