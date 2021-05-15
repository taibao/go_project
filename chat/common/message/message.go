package message

const unreg_code = 500
const succ_code = 200

const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

type Message struct{
	Type string  `json:"type"`//消息类型
	Data string `json:"data"` //消息数据
}

type LoginMes struct{
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct{
	Code int  `json:"code"`//返回消息 状态码 500表示未注册，200成功
	Error string  `json:"error"`//错误信息
}

type RegisterMes struct{
	User User `json:"user"` //类型就是User结构体
}

type RegisterResMes struct{
	Code int `json:"code"` //返回状态码400 表示该用户已经占有，200表示注册成功
	Error string `json:"error"` //返回错误信息
}



