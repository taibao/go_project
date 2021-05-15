package processes

import (
	"chat/common/message"
	"chat/server/utils"
	"fmt"
	"io"
	"net"
)

//创建Processor的结构体
type Processor struct{
	Conn  net.Conn
}

//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) ServerProcessMes(mes *message.Message) (err error){

	//判断是否接受到群发消息
	fmt.Println("mes=",mes)

	switch mes.Type {
	case message.LoginMesType:
		//处理登录逻辑
		processes := &UserProcess{
			Conn: this.Conn,
		}
		err = processes.ServerProcessLogin(mes)

	case message.RegisterMesType:
		//处理注册逻辑
		processes := &UserProcess{
			Conn: this.Conn,
		}
		err = processes.ServerProcessRegister(mes)
	case message.SmsMesType:
		//创建SmsProcess实例完成转发群聊消息
		smsProcess := &SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理。。。")
	}
	return
}

func (this *Processor) Process2() (err error){
	//客户端循环发送信息
	for{
		//这里我们将读取数据包，直接封装成一个函数readPkg，返回Message，Err
		//创建一个Tranfer实例完成读包任务
		tf := &utils.Transfer{
			Conn:this.Conn,
		}

		mes,err := tf.ReadPkg()
		if err != nil{
			if err ==io.EOF{
				fmt.Println("客户端退出，服务器端也退出")
				return err
			}else{
				fmt.Println("readpkg err=",err)
				return err
			}
		}

		fmt.Println("发送内容=",mes)

		err = this.ServerProcessMes(&mes)
		if err != nil{
			return err
		}
	}
}