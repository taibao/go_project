package processes

import (
	"chat/common/message"
	"chat/server/utils"
	"encoding/json"
	"fmt"
	"net"
	"chat/server/model"
)

type UserProcess struct{
	//字段
	Conn net.Conn
}



//处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message)(err error){
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

	//需要到redis数据库去完成验证
	//使用model.MyUserDao 到redis去验证
	user,err := model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	if err != nil{
		loginResMes.Code = 500
		//

	}else{
		loginResMes.Code = 200
		fmt.Println(user,"登录成功")
	}

	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456"{
	//	//合法
	//	loginResMes.Code = 200
	//}else{
	//	//不合法
	//	loginResMes.Code = 500
	//	loginResMes.Error = "该用户不存在，请注册再使用..."
	//}

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
	//因为使用分层模式（mvc），我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
