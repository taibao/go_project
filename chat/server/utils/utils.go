package utils

import (
	"chat/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

//将方法关联到结构体中
type Transfer struct{
	//分析它应该有哪些字段
	Conn net.Conn
	Buf [8096]byte //这时传输时，使用缓冲

}

func (this *Transfer) ReadPkg() (mes message.Message,err error){
	//conn.Read在conn没有被关闭的情况下才会阻塞
	//如果客户端关闭了conn，就不会阻塞了
	_,err = this.Conn.Read(this.Buf[:4])
	if err != nil{
		err = errors.New("read pkg header error")
		return
	}
	//fmt.Println("读到的buf=",buf[:4])
	//根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	//根据pkgLen读取消息内容
	n,err := this.Conn.Read(this.Buf[:pkgLen]) //是把conn的数据读到buf里面去
	if n != int(pkgLen) || err != nil{
		//err = errors.New("read pkg body error")
		return
	}

	//把pkgLen反序列化成-》message.Message
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)
	if err != nil{
		fmt.Println("json.unmarshal err =",err,mes)
		return
	}

	return
}

//连接数据和判断写入包
func (this *Transfer) WritePkg(data []byte) (err error){
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[0:4],pkgLen)
	//发送长度
	n,err := this.Conn.Write(this.Buf[:4])
	if n!=4 || err !=nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return err
	}

	//发送data本身
	n,err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}
	return
}

