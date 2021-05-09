package message

const unreg_code = 500
const succ_code = 200

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
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
