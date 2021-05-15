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

func (this *UserProcess)  ServerProcessRegister(mes *message.Message) (err error){
	//1.先从mes中取出mes.Data,并直接反序列化成RegisterMes
	var registerMes message.RegisterMes //注册数据
	err = json.Unmarshal([]byte(mes.Data),&registerMes)
	if err != nil{
		fmt.Println("json.Marshal fail err=",err)
		return
	}

	//注册消息
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	//在数据库中验证数据是否已存在
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil{
		if err == model.ERROR_USER_EXISTS{
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		}else{
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	}else{
		registerResMes.Code = 200
	}

	//序列化
	data,err := json.Marshal(registerResMes)
	if err != nil{
		fmt.Println("json.Marshal fail",err)
		return
	}
	//将data赋值给resMes
	//将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes 进行序列化，准备发送
	data,err = json.Marshal(resMes)
	if err != nil{
		fmt.Println("json.Marshal fail",err)
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
		if err == model.ERROR_USER_NOTEXISTS{
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		}else if err == model.ERROR_USER_PWD{
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		}else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误...."
		}
	}else{
		loginResMes.Code = 200
		fmt.Println(user,"登录成功")
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
	//因为使用分层模式（mvc），我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
