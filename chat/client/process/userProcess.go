package process

import (
	"chat/client/utils"
	"chat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct{

}

func (this *UserProcess) Register(userId int,userPwd string,userName string) (err error) {
	//连接到服务器端
	conn,err :=  net.Dial("tcp","localhost:8889")
	if err != nil{
		fmt.Println("net.Dial err =",err)
		return
	}
	//延时关闭
	defer conn.Close()

	//准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3:创建一个RegisterMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4:将RegisterMes序列化
	data,err := json.Marshal(registerMes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	//5:把data赋值给mes.data字段
	mes.Data = string(data)
	//6:将mes进行序列化
	data,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.marshal err=",err)
		return
	}

	//还需要处理服务器端返回的消息
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn:conn,
	}

	//发送data给服务器端
	err = tf.WritePkg(data)
	if err !=nil{
		fmt.Println("注册发送信息错误 err=",err)
		return
	}

	mes,err = tf.ReadPkg() //mes 就是
	if err != nil{
		fmt.Println("readPkg(conn) err=",err)
		return
	}

	//将mes的Data部分反序列化成RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	if registerResMes.Code == 200{
		fmt.Println("注册成功，请重新登录")
	}else {
		fmt.Println( registerResMes.Error)
	}
	os.Exit(0)
	return
}

//写一个函数，完成登录
func (this *UserProcess) Login(userId int,userPwd string) {
	//fmt.Printf("userId=%d userPwd=%v \n",userId,userPwd)
	//return nil

	//连接到服务器端
	conn,err :=  net.Dial("tcp","localhost:8889")
	if err != nil{
		fmt.Println("net.Dial err =",err)
		return
	}
	//延时关闭
	defer conn.Close()

	//准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType

	//3:创建一个loginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4:将loginMes序列化
	data,err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	//把data赋值给mes.data字段
	mes.Data = string(data)
	//将mes进行序列化
	data,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.marshal err",err)
		return
	}

	//此时data为我们要发送的消息
	//先把data的长度发送给服务器
	//先获取到data的长度--》转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
	//发送长度
	n,err := conn.Write(buf[:4])
	if n!=4 || err !=nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}

	fmt.Printf("客户端，发送消息的长度=%d 内容=%s\n",len(data),string(data))

	//发送消息本身
	_,err = conn.Write(data)
	if err != nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}

	//休眠20
	//time.Sleep(20*time.Second)
	//fmt.Println("休眠了20...")
	//还需要处理服务器端返回的消息
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn:conn,
	}

	mes,err = tf.ReadPkg() //mes

	fmt.Println("查看mes返回",mes)

	if err != nil{
		fmt.Println("ReadPkg(conn) failed err=",err)
		return
	}

	//将mes的Data部分反序列化成loginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200{
		fmt.Println("登录成功")
		//显示当前在线用户列表，遍历loginResMes.UsersId
		fmt.Println("当前在线用户列表如下：")
		for _,v := range loginResMes.UsersId{
			fmt.Println("用户id:\t",v)
		}
		fmt.Println()

		//这里我们还需要客户端启动协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则接受并显示在客户端的终端

		go serverProcessMes(conn)
		//1. 显示登录成功的菜单[循环显示]
		for{
			ShowMenu()
		}

	}else {
		fmt.Printf("登录失败，错误信息=%v\n",loginResMes.Error)
	}

	return
}