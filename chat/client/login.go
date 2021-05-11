package client

import (
	"chat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//写一个函数，完成登录
func Login(userId int,userPwd string) (err error){
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

	//将loginMes序列化
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
		return err
	}

	//fmt.Printf("客户端，发送消息的长度=%d 内容=%s",len(data),string(data))

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
	mes,err = ReadPkg(conn) //mes
	if err != nil{
		fmt.Println("ReadPkg(conn) failed err=",err)
		return
	}

	//将mes的Data部分反序列化成loginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200{
		fmt.Println("登录成功")
	}else if loginResMes.Code == 500{
		fmt.Println("登录失败，错误信息=",loginResMes.Error)
	}

	return
}