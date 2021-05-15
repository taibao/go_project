package processes

import (
	"chat/common/message"
	"chat/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct{
	//
}

//写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message){
	//遍历服务器端的onlineUsers map[int] *UserProcess,
	//将消息转发取出
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err !=nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}

	data,err := json.Marshal(mes)
	if err!=nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	for id, up := range userMgr.onlineUsers{
		if id == smsMes.UserId{
			continue
		}
		this.SendMesToEachOnlineUser(data,up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte,conn net.Conn){
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn:conn,
	}

	//发送data给服务器端
	err := tf.WritePkg(data)
	if err !=nil{
		fmt.Println("转发消息失败 err=",err)
		return
	}
}