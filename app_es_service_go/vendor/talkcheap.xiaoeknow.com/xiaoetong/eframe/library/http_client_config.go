package library

type HttpClientConfig struct {
	Name string //名称
	DialTimeoutSecond int // 连接超时
	DialKeepAliveSecond int //开启长连接
	MaxIdleConnections int //最大空闲连接数
	MaxIdleConnectionsPerHost int //单Host最大空闲连接数
	IdleConnTimeoutSecond int // 空闲连接超时
}

