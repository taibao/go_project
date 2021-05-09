package main

import (
	"fmt"
	"net" //做网络socket开发时，net包含我们需要的所有方法和函数
)

//处理客户端数据
func process(conn net.Conn){
	//循环接收客户端发送的数据
	defer conn.Close() //关闭conn
	for{
		//创建一个新的切片
		buf := make([]byte,1024)
		//conn.Read(buf)
		//1： 等待客户端通过conn发送消息
		//2: 如果客户端没有write【发送】，那么协程就阻塞在此
		//fmt.Println("服务器在等待客户端发送信息",conn.RemoteAddr().String())
		n,err := conn.Read(buf) //从conn读取
		if err != nil{
			fmt.Printf("客户端已退出 err=%v \n",err)
			return
		}
		//显示客户端发送的内容到服务器的终端
		fmt.Println(string(buf[:n]))
	}
}


func main(){
	fmt.Println("服务器开始监听。。。。")
	//1. tcp表示使用网络协议是tcp
	//2. 0.0.0.0:8888 表示在本地监听 8888端口
	listen,err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil{
		fmt.Println("listen err=",err)
		return
	}

	defer listen.Close() //close关闭该接口，并使任何阻塞的Accept操作都不会再阻塞并返回错误

	//循环等待客户端来连接
	for{
		//等待客户端来连接
		fmt.Println("等待客户端来连接...")
		conn,err := listen.Accept() //连接到conn
		if err != nil{
			fmt.Println("Accept() err=",err)
		}else{
			//输出ip端口
			fmt.Printf("Accept() suc con=%v 客户端ip=%v \n",conn,conn.RemoteAddr().String())
		}
		//起协程为客户端服务
		go process(conn)
	}

	//fmt.Println(listen)
}